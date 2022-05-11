package main

import (
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"

	"github.com/jalavosus/pixelimagedl"
)

var listCmd = cli.Command{
	Name: "list",
	Flags: []cli.Flag{
		&downloadTypeFlag,
		&deviceNameFlag,
	},
	Action: listCmdAction,
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
