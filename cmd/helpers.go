package main

import (
	"path/filepath"
	"strings"

	"github.com/jalavosus/pixelimagedl"
	"github.com/jalavosus/pixelimagedl/internal"
)

func absPath() string {
	p, _ := filepath.Abs("./")
	return p
}

func makeSmallDeviceName(deviceName string) string {
	return strings.ToLower(strings.Replace(deviceName, " ", "", -1))
}

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
