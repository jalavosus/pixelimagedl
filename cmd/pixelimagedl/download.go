package main

import (
	"context"

	"github.com/urfave/cli/v2"

	"github.com/jalavosus/pixelimagedl/pkg/pixelimagedl"
)

var downloadCmd = cli.Command{
	Name: "download",
	Flags: []cli.Flag{
		&downloadTypeFlag,
		&deviceNameFlag,
		&downloadTimeoutFlag,
		&outDirFlag,
	},
	Action: WithFlags(downloadCmdAction),
}

func downloadCmdAction(c *cli.Context, parsedFlags ParsedFlags) error {
	ctx, cancel := context.WithTimeout(c.Context, parsedFlags.DownloadTimeout)
	defer cancel()

	return pixelimagedl.DownloadLatest(
		ctx,
		parsedFlags.Device,
		parsedFlags.DownloadType,
		parsedFlags.OutDir,
	)
}
