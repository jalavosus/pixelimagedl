package pixelimagedl

//go:generate stringer -type Pixel,Codename,DownloadType -linecomment -output enums_string.go

type (
	Pixel        uint
	Codename     uint
	DownloadType uint8
)

const (
	Factory DownloadType = iota // factory
	OTA                         // OTA
)

const (
	UnknownDevice Pixel = iota // unknown
	Pixel4                     // Pixel 4
	Pixel4XL                   // Pixel 4 XL
	Pixel4a                    // Pixel 4a
	Pixel4a5G                  // Pixel 4a 5G
	Pixel5                     // Pixel 5
	Pixel5a                    // Pixel 5a
	Pixel6                     // Pixel 6
	Pixel6Pro                  // Pixel 6 Pro
)

var AllPixelNames = []Pixel{
	Pixel4,
	Pixel4XL,
	Pixel4a,
	Pixel4a5G,
	Pixel5,
	Pixel5a,
	Pixel6,
	Pixel6Pro,
}

const (
	UnknownCodename Codename = iota // unknown
	Flame                           // flame
	Coral                           // coral
	Sunfish                         // sunfish
	Bramble                         // bramble
	Redfin                          // redfin
	Barbet                          // barbet
	Oriole                          // oriole
	Raven                           // raven
)

var AllCodenames = []Codename{
	Flame,
	Coral,
	Sunfish,
	Bramble,
	Redfin,
	Barbet,
	Oriole,
	Raven,
}

var deviceCodenameMap = map[Pixel]Codename{
	Pixel4:    Flame,
	Pixel4XL:  Coral,
	Pixel4a:   Sunfish,
	Pixel4a5G: Bramble,
	Pixel5:    Redfin,
	Pixel5a:   Barbet,
	Pixel6:    Oriole,
	Pixel6Pro: Raven,
}

func DeviceFromCodename(c Codename) Pixel {
	for k, v := range deviceCodenameMap {
		if v == c {
			return k
		}
	}

	return UnknownDevice
}

var AllDeviceNames = []string{
	Pixel4.String(),
	Pixel4XL.String(),
	Pixel4a.String(),
	Pixel4a5G.String(),
	Pixel5.String(),
	Pixel5a.String(),
	Pixel6.String(),
	Pixel6Pro.String(),
	Flame.String(),
	Coral.String(),
	Sunfish.String(),
	Bramble.String(),
	Redfin.String(),
	Barbet.String(),
	Oriole.String(),
	Raven.String(),
}
