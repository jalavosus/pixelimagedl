package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"

	"github.com/jalavosus/pixelimagedl"
	"github.com/jalavosus/pixelimagedl/internal"
)

func absPath() string {
	p, _ := filepath.Abs("./")
	return p
}

var (
	downloadTypeFlag = cli.StringFlag{
		Name:     "imagetype",
		Usage:    "`type` of image zip to download (factory or ota)",
		Aliases:  []string{"t"},
		Required: false,
		Value:    pixelimagedl.Factory.String(),
	}
	deviceNameFlag = cli.StringFlag{
		Name:     "device",
		Usage:    "`name` (or codename) of the device to download an image for",
		Aliases:  []string{"d"},
		Required: true,
	}
	outDirFlag = cli.PathFlag{
		Name:     "outdir",
		Usage:    "`dir`ectory to place downloaded image file in",
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

func makeSmallDeviceName(deviceName string) string {
	return strings.ToLower(strings.Replace(deviceName, " ", "", -1))
}

func validateDevice(raw string) (pixelimagedl.Pixel, bool) {
	raw = strings.ToLower(raw)
	for _, p := range pixelimagedl.AllPixelNames {
		if makeSmallDeviceName(p.String()) == raw {
			return p, true
		}
	}

	for _, c := range pixelimagedl.AllCodenames {
		if makeSmallDeviceName(c.String()) == raw {
			return pixelimagedl.DeviceFromCodename(c), true
		}
	}

	return pixelimagedl.UnknownDevice, false
}

var allowedDeviceNames = internal.Map(pixelimagedl.AllDeviceNames, makeSmallDeviceName)

func downloadCmdAction(c *cli.Context) error {
	rawDeviceName := deviceNameFlag.Get(c)
	rawImageKind := downloadTypeFlag.Get(c)

	deviceName, ok := validateDevice(rawDeviceName)
	if !ok {
		return errors.Errorf("invalid device name %[1]s. Allowed values: %[2]s", rawDeviceName, strings.Join(allowedDeviceNames, ", "))
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
