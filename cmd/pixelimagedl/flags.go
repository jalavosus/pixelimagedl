package main

import (
	"context"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v3"

	"github.com/jalavosus/pixelimagedl/pkg/pixelimagedl"
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
	outDirFlag = cli.StringFlag{
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

func WithFlags(fn func(context.Context, ParsedFlags) error) cli.ActionFunc {
	return func(ctx context.Context, command *cli.Command) error {
		parsedFlags, flagsErr := parseFlags(command)
		if flagsErr != nil {
			return flagsErr
		}

		return fn(ctx, parsedFlags)
	}
}

func parseFlags(cmd *cli.Command) (ParsedFlags, error) {
	var parsedFlags ParsedFlags

	rawDeviceName := cmd.String(deviceNameFlag.Name)
	rawImageKind := cmd.String(downloadTypeFlag.Name)

	deviceName, ok := validateDevice(rawDeviceName)
	if !ok {
		return parsedFlags, errors.Errorf("invalid device name %[1]s. Allowed values: %[2]s", rawDeviceName, strings.Join(allowedDeviceNames, ", "))
	}

	downloadKind, ok := validateImageKind(rawImageKind)
	if !ok {
		return parsedFlags, errors.Errorf("invalid download kind %[1]s. Allowed values: %[2]s", rawImageKind, strings.Join(allowedDownloadTypes, ", "))
	}

	downloadTimeout := cmd.Duration(downloadTimeoutFlag.Name)
	outDir := cmd.String(outDirFlag.Name)

	parsedFlags = ParsedFlags{
		Device:          deviceName,
		DownloadType:    downloadKind,
		DownloadTimeout: downloadTimeout,
		OutDir:          outDir,
	}

	return parsedFlags, nil
}
