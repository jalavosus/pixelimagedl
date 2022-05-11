package pixelimagedl

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"

	"github.com/jalavosus/pixelimagedl/pkg/download"
)

func DownloadLatest(device Pixel, downloadType DownloadType, outDir string) error {
	images, err := ScrapeDeviceImages(device, downloadType)
	if err != nil {
		err = errors.WithMessagef(err, "error scraping available %[1]s images for device %[2]s", downloadType.String(), device.String())
		return err
	}

	latest := last(images)

	log.Printf("latest stable %[1]s image for %[2]s is %[3]s (%[4]s)\n", downloadType.String(), device.String(), latest.Version, latest.BuildNumber)

	var filename string

	downloadUri := latest.DownloadURI
	split := strings.Split(downloadUri, "/")

	filename = last(split)
	if !filepath.IsAbs(outDir) {
		outDir, err = filepath.Abs(outDir)
		if err != nil {
			return err
		}
	}
	filename = filepath.Join(outDir, filename)

	log.Printf("downloading %[1]s image from %[2]s\n", downloadType.String(), downloadUri)

	resp, err := http.Get(downloadUri)
	if err != nil {
		err = errors.WithMessagef(err, "error downloading file at url %[1]s", downloadUri)
		return err
	}

	log.Printf("saving %[1]s image to %[2]s\n", downloadType.String(), filename)
	numBytes, err := download.ReadData(resp, filename, dlBufSize())
	if err != nil {
		return err
	}

	log.Printf("saved %-.1[1]fGb to %[2]s", download.GbFromBytes(numBytes), filename)

	gotSha, shaMatch := checkSha(filename, latest.SHA256Sum)
	if !shaMatch {
		return errors.Errorf("SHA256 mismatch; expected %[1]s, sum of downloaded file is %[2]s", latest.SHA256Sum, gotSha)
	} else {
		log.Printf("SHA256 sum %[1]s of downloaded file matches expected\n", gotSha)
	}

	return nil
}

func checkSha(filename, wantSha string) (string, bool) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = f.Close()
	}()

	h := sha256.New()
	if _, err = io.Copy(h, f); err != nil {
		panic(err)
	}

	check := fmt.Sprintf("%x", h.Sum(nil))

	return check, check == wantSha
}
