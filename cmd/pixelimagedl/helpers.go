package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/jalavosus/pixelimagedl/internal"
	"github.com/jalavosus/pixelimagedl/pkg/pixelimagedl"
)

func absPath() string {
	p, _ := filepath.Abs("./")
	return p
}

func makeCliFlagVal(flagVal string) string {
	return strings.ToLower(strings.ReplaceAll(flagVal, " ", ""))
}

func validateImageKind(raw string) (pixelimagedl.DownloadType, bool) {
	return checkFlagVal(allDownloadTypes, raw)
}

func validateDevice(raw string) (pixelimagedl.Pixel, bool) {
	if pixelName, isPixelName := checkFlagVal(pixelimagedl.AllPixelNames, raw); isPixelName {
		return pixelName, true
	}

	if codename, isCodename := checkFlagVal(pixelimagedl.AllCodenames, raw); isCodename {
		return pixelimagedl.DeviceFromCodename(codename), true
	}

	return pixelimagedl.UnknownDevice, false
}

func checkFlagVal[T fmt.Stringer](allowedVals []T, check string) (T, bool) {
	var emptyVal T

	check = strings.ToLower(check)

	for _, val := range allowedVals {
		if makeCliFlagVal(val.String()) == check {
			return val, true
		}
	}

	return emptyVal, false
}

var (
	allDownloadTypes     = []pixelimagedl.DownloadType{pixelimagedl.Factory, pixelimagedl.OTA}
	allowedDeviceNames   = internal.Map(pixelimagedl.AllDeviceNames, makeCliFlagVal)
	allowedDownloadTypes = internal.Map(
		[]string{pixelimagedl.Factory.String(), pixelimagedl.OTA.String()},
		makeCliFlagVal,
	)
)
