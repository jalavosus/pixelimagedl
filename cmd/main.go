package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"

	"github.com/jalavosus/pixelimagedl"
)

func absPath() string {
	p, _ := filepath.Abs("./")
	return p
}

var (
	downloadTypeFlag = cli.StringFlag{
		Name:     "type",
		Aliases:  []string{"t"},
		Required: false,
		Value:    pixelimagedl.Factory.String(),
	}
	deviceNameFlag = cli.StringFlag{
		Name:     "device",
		Aliases:  []string{"d"},
		Required: true,
	}
	outDirFlag = cli.PathFlag{
		Name:     "outDir",
		Aliases:  []string{"o"},
		Required: false,
		Value:    absPath(),
	}
)

var (
	downloadCmd = cli.Command{
		Name: "download",
		Flags: []cli.Flag{
			&downloadTypeFlag,
			&deviceNameFlag,
			&outDirFlag,
		},
		Action: downloadCmdAction,
	}
	listCmd = cli.Command{
		Name: "list",
		Flags: []cli.Flag{
			&downloadTypeFlag,
			&deviceNameFlag,
		},
		Action: listCmdAction,
	}
)

func validateImageKind(raw string) (pixelimagedl.DownloadType, bool) {
	switch strings.ToLower(raw) {
	case strings.ToLower(pixelimagedl.Factory.String()):
		return pixelimagedl.Factory, true
	case strings.ToLower(pixelimagedl.OTA.String()):
		return pixelimagedl.OTA, true
	default:
		return pixelimagedl.Factory, true
	}
}

type DeviceType interface {
	pixelimagedl.Pixel | pixelimagedl.Codename
	String() string
}

func makeSmallDeviceName[T DeviceType](deviceName T) string {
	return strings.ToLower(strings.Replace(deviceName.String(), " ", "", -1))
}

func validateDevice(raw string) (pixelimagedl.Pixel, bool) {
	switch strings.ToLower(raw) {
	case makeSmallDeviceName(pixelimagedl.Pixel4):
		return pixelimagedl.Pixel4, true
	case makeSmallDeviceName(pixelimagedl.Pixel4XL):
		return pixelimagedl.Pixel4XL, true
	case makeSmallDeviceName(pixelimagedl.Pixel4a):
		return pixelimagedl.Pixel4a, true
	case makeSmallDeviceName(pixelimagedl.Pixel4a5G):
		return pixelimagedl.Pixel4a5G, true
	case makeSmallDeviceName(pixelimagedl.Pixel5):
		return pixelimagedl.Pixel5, true
	case makeSmallDeviceName(pixelimagedl.Pixel5a):
		return pixelimagedl.Pixel5a, true
	case makeSmallDeviceName(pixelimagedl.Pixel6):
		return pixelimagedl.Pixel6, true
	case makeSmallDeviceName(pixelimagedl.Pixel6Pro):
		return pixelimagedl.Pixel6Pro, true
	case makeSmallDeviceName(pixelimagedl.Flame):
		return pixelimagedl.Pixel4, true
	case makeSmallDeviceName(pixelimagedl.Coral):
		return pixelimagedl.Pixel4XL, true
	case makeSmallDeviceName(pixelimagedl.Sunfish):
		return pixelimagedl.Pixel4a, true
	case makeSmallDeviceName(pixelimagedl.Bramble):
		return pixelimagedl.Pixel4a5G, true
	case makeSmallDeviceName(pixelimagedl.Redfin):
		return pixelimagedl.Pixel5, true
	case makeSmallDeviceName(pixelimagedl.Barbet):
		return pixelimagedl.Pixel5a, true
	case makeSmallDeviceName(pixelimagedl.Oriole):
		return pixelimagedl.Pixel6, true
	case makeSmallDeviceName(pixelimagedl.Raven):
		return pixelimagedl.Pixel6Pro, true
	default:
		return pixelimagedl.Unknown, false
	}
}

var allowedDeviceNames = []string{
	makeSmallDeviceName(pixelimagedl.Pixel4),
	makeSmallDeviceName(pixelimagedl.Pixel4XL),
	makeSmallDeviceName(pixelimagedl.Pixel4a),
	makeSmallDeviceName(pixelimagedl.Pixel4a5G),
	makeSmallDeviceName(pixelimagedl.Pixel5),
	makeSmallDeviceName(pixelimagedl.Pixel5a),
	makeSmallDeviceName(pixelimagedl.Pixel6),
	makeSmallDeviceName(pixelimagedl.Pixel6Pro),
	makeSmallDeviceName(pixelimagedl.Flame),
	makeSmallDeviceName(pixelimagedl.Coral),
	makeSmallDeviceName(pixelimagedl.Sunfish),
	makeSmallDeviceName(pixelimagedl.Bramble),
	makeSmallDeviceName(pixelimagedl.Redfin),
	makeSmallDeviceName(pixelimagedl.Barbet),
	makeSmallDeviceName(pixelimagedl.Oriole),
	makeSmallDeviceName(pixelimagedl.Raven),
}

func downloadCmdAction(c *cli.Context) error {
	rawDeviceName := deviceNameFlag.Get(c)
	rawImageKind := downloadTypeFlag.Get(c)

	deviceName, ok := validateDevice(rawDeviceName)
	if !ok {
		return errors.Errorf("invalid device name %[1]s.\nAllowed values: %[2]s", rawDeviceName, strings.Join(allowedDeviceNames, ", "))
	}

	downloadKind, ok := validateImageKind(rawImageKind)
	if !ok {
		return errors.Errorf("invalid download kind %[1]s", rawImageKind)
	}

	err := pixelimagedl.DownloadLatest(deviceName, downloadKind, outDirFlag.Get(c))
	if err != nil {
		return err
	}

	return nil
}

func listCmdAction(c *cli.Context) error {
	rawDeviceName := deviceNameFlag.Get(c)
	rawImageKind := downloadTypeFlag.Get(c)

	deviceName, ok := validateDevice(rawDeviceName)
	if !ok {
		return errors.Errorf("invalid device name %[1]s", rawDeviceName)
	}

	downloadKind, ok := validateImageKind(rawImageKind)
	if !ok {
		return errors.Errorf("invalid download kind %[1]s", rawImageKind)
	}

	data, err := pixelimagedl.ScrapeDeviceImages(deviceName, downloadKind)
	if err != nil {
		return err
	}

	for _, d := range data {
		d.PrettyPrint()
	}

	return nil
}

func main() {
	app := &cli.App{
		Name: "pixel image downloader",
		Authors: []*cli.Author{
			{Name: "jalavosus", Email: "alavosus.james@gmail.com"},
		},
		Commands: []*cli.Command{
			&downloadCmd,
			&listCmd,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
