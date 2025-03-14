package main

import (
	"context"
	"time"

	"github.com/urfave/cli/v3"

	"github.com/jalavosus/pixelimagedl/pkg/pixelimagedl"
)

var listCmd = cli.Command{
	Name: "list",
	Flags: []cli.Flag{
		&downloadTypeFlag,
		&deviceNameFlag,
	},
	Action: WithFlags(listCmdAction),
}

func listCmdAction(ctx context.Context, parsedFlags ParsedFlags) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	data, err := pixelimagedl.ListDeviceImages(ctx, parsedFlags.Device, parsedFlags.DownloadType)
	if err != nil {
		return err
	}

	for _, d := range data {
		d.PrettyPrint()
	}

	return nil
}
