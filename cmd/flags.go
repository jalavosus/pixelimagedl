package main

import (
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"

	"github.com/jalavosus/pixelimagedl"
)

var (
	downloadTypeFlag = cli.StringFlag{
		Name:     "imagetype",
		Usage:    "`type` of image zip to download (factory or ota)",
		Aliases:  []string{"t", "type"},
		Required: false,
		Value:    pixelimagedl.Factory.String(),
	}
	deviceNameFlag = cli.StringFlag{
		Name:     "device",
		Usage:    "`name` (or codename) of the device to download an image for",
		Aliases:  []string{"d"},
		Required: true,
	}
	downloadTimeoutFlag = cli.DurationFlag{
		Name:     "timeout",
		Usage:    "`timeout` for file downloads",
		Value:    15 * time.Minute,
		Required: false,
	}
	outDirFlag = cli.PathFlag{
		Name:     "outdir",
		Usage:    "`dir`ectory to place downloaded image file in",
		Aliases:  []string{"o"},
		Required: false,
		Value:    absPath(),
	}
)

type ParsedFlags struct {
	OutDir          string
	Device          pixelimagedl.Pixel
	DownloadTimeout time.Duration
	DownloadType    pixelimagedl.DownloadType
}

func WithFlags(fn func(*cli.Context, ParsedFlags) error) cli.ActionFunc {
	return func(c *cli.Context) error {
		parsedFlags, flagsErr := parseFlags(c)
		if flagsErr != nil {
			return flagsErr
		}

		return fn(c, parsedFlags)
	}
}

func parseFlags(c *cli.Context) (ParsedFlags, error) {
	var parsedFlags ParsedFlags

	rawDeviceName := deviceNameFlag.Get(c)
	rawImageKind := downloadTypeFlag.Get(c)

	deviceName, ok := validateDevice(rawDeviceName)
	if !ok {
		return parsedFlags, errors.Errorf("invalid device name %[1]s. Allowed values: %[2]s", rawDeviceName, strings.Join(allowedDeviceNames, ", "))
	}

	downloadKind, ok := validateImageKind(rawImageKind)
	if !ok {
		return parsedFlags, errors.Errorf("invalid download kind %[1]s. Allowed values: %[2]s", rawImageKind, strings.Join(allowedDownloadTypes, ", "))
	}

	downloadTimeout := downloadTimeoutFlag.Get(c)
	outDir := outDirFlag.Get(c)

	parsedFlags = ParsedFlags{
		Device:          deviceName,
		DownloadType:    downloadKind,
		DownloadTimeout: downloadTimeout,
		OutDir:          outDir,
	}

	return parsedFlags, nil
}
