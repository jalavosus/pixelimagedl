package main

import (
	"github.com/urfave/cli/v2"

	"github.com/jalavosus/pixelimagedl"
)

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
