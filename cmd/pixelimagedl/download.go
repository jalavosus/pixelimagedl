package main

import (
	"context"

	cli "github.com/urfave/cli/v3"

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

func downloadCmdAction(ctx context.Context, parsedFlags ParsedFlags) error {
	ctx, cancel := context.WithTimeout(ctx, parsedFlags.DownloadTimeout)
	defer cancel()

	return pixelimagedl.DownloadLatest(
		ctx,
		parsedFlags.Device,
		parsedFlags.DownloadType,
		parsedFlags.OutDir,
	)
}
