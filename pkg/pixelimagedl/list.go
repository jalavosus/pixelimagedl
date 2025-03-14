package pixelimagedl

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"sort"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"

	"github.com/jalavosus/pixelimagedl/internal"
)

func ListDeviceImages(ctx context.Context, device Pixel, downloadType DownloadType) ([]PixelImage, error) {
	codename := deviceCodenameMap[device]

	data, err := scrapeData(ctx, codename, downloadType)
	if err != nil {
		return nil, err
	}

	data = sortDataSlice(data)

	return data, nil
}

func scrapeData(ctx context.Context, codename Codename, downloadType DownloadType) ([]PixelImage, error) {
	var (
		deviceImages []PixelImage
		downloadUri  string
		cookieData   string
	)

	switch downloadType {
	case Factory:
		downloadUri = internal.StableFactoryImagesURL
		cookieData = internal.FactoryAcksCookie
	case OTA:
		downloadUri = internal.StableOTAImagesURL
		cookieData = internal.OTAAcksCookie
	}

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, downloadUri, http.NoBody)
	req.Header.Set("cookie", cookieData)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		err = errors.WithMessagef(err, "error requesting url %[1]s", downloadUri)
		return nil, err
	}

	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			log.Printf("error closing file download body: %v\n", closeErr)
		}
	}()

	pageBody, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		err = errors.WithMessage(err, "error reading response body")
		return nil, err
	}

	deviceTable := findDeviceTable(codename, pageBody)
	if deviceTable == nil {
		fmt.Println(pageBody.Text())
		return nil, nil
	}

	deviceImages = parseRows(deviceTable, downloadType)

	return deviceImages, nil
}

func findDeviceTable(codename Codename, pageBody *goquery.Document) *goquery.Selection {
	var (
		foundHeader *goquery.Selection
		deviceTable *goquery.Selection
	)

	headers := pageBody.Find("h2")
	headers.Each(func(idx int, s *goquery.Selection) {
		elemId, ok := s.Attr("id")
		if ok && elemId == codename.String() {
			foundHeader = s
		}
	})

	if foundHeader != nil {
		deviceTable = foundHeader.Next().Find("tbody")
	}

	return deviceTable
}

func parseRows(tableBody *goquery.Selection, downloadType DownloadType) []PixelImage {
	var parsed []PixelImage

	tableBody.Find("tr").Each(func(idx int, s *goquery.Selection) {
		imageData := PixelImage{}

		rowData := s.Find("td")

		fullBuild := rowData.First()
		fullBuildText := fullBuild.Text()

		imageData.Version,
			imageData.BuildNumber,
			imageData.BuildDate,
			imageData.BuildComment = parseVersionString(fullBuildText)

		linkData := fullBuild.Next()
		if downloadType == Factory {
			linkData = linkData.Next()
		}

		downloadLink, ok := linkData.Find("a").Attr("href")
		if ok {
			imageData.DownloadURI = downloadLink
		}

		imageData.SHA256Sum = linkData.Next().Text()

		parsed = append(parsed, imageData)
	})

	return parsed
}

var (
	versionRegex = regexp.MustCompile(`(\d{1,2}\.\d{1,2}\.\d*)`)
	buildRegex   = regexp.MustCompile(`\((.*)\)`)
)

func parseVersionString(buildNum string) (version, buildNumber, buildDate, buildComment string) {
	version = versionRegex.FindString(buildNum)

	b := buildRegex.FindString(buildNum)
	b = strings.TrimPrefix(b, "(")
	b = strings.TrimSuffix(b, ")")
	bSplit := strings.Split(b, ",")

	buildNumber = strings.TrimSpace(bSplit[0])
	buildDate = strings.TrimSpace(bSplit[1])
	if len(bSplit) > 2 {
		buildComment = strings.TrimSpace(strings.Join(bSplit[2:], ", "))
		buildComment = strings.ReplaceAll(buildComment, "  ", " ")
	}

	return
}

func sortDataSlice(data []PixelImage) []PixelImage {
	sort.Slice(data, func(i, j int) bool {
		majorI, minorI, extraI := getBuildMajorMinor(data[i].BuildNumber)
		majorJ, minorJ, extraJ := getBuildMajorMinor(data[j].BuildNumber)

		if majorI == majorJ {
			if minorI == minorJ {
				return extraI < extraJ
			}

			return minorI < minorJ
		}

		return majorI < majorJ
	})

	return data
}

func getBuildMajorMinor(buildNumber string) (major, minor int64, extra string) {
	split := strings.Split(buildNumber, ".")
	majorStr := split[1]
	minorStr := split[2]
	if len(split) == 4 {
		extra = split[3]
	}

	major = internal.ParseInt64(majorStr)
	minor = internal.ParseInt64(minorStr)

	return
}
