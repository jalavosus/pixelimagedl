package main

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"

	"github.com/jalavosus/pixelimagedl"
)

var downloadCmd = cli.Command{
	Name: "download",
	Flags: []cli.Flag{
		&downloadTypeFlag,
		&deviceNameFlag,
		&outDirFlag,
	},
	Action: downloadCmdAction,
}

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
