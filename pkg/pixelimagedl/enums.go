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
	Pixel6a                    // Pixel 6a
	Pixel7                     // Pixel 7
	Pixel7Pro                  // Pixel 7 Pro
	Pixel7a                    // Pixel 7a
	PixelTablet                // Pixel Tablet
	PixelFold                  // Pixel Fold
	Pixel8                     // Pixel 8
	Pixel8Pro                  // Pixel 8 Pro
	Pixel8a                    // Pixel 8a
	Pixel9                     // Pixel 9
	Pixel9Pro                  // Pixel 9 Pro
	Pixel9ProXL                // Pixel 9 Pro XL
	Pixel9ProFold              // Pixel 9 Pro Fold
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
	Pixel6a,
	Pixel7,
	Pixel7Pro,
	Pixel7a,
	PixelTablet,
	PixelFold,
	Pixel8,
	Pixel8Pro,
	Pixel8a,
	Pixel9,
	Pixel9Pro,
	Pixel9ProXL,
	Pixel9ProFold,
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
	Bluejay                         // bluejay
	Panther                         // panther
	Cheetah                         // cheetah
	Lynx                            // lynx
	TangorPro                       // tangorpro
	Felix                           // felix
	Shiba                           // shiba
	Husky                           // husky
	Akita                           // akita
	Tokay                           // tokay
	Caiman                          // caiman
	Komodo                          // komodo
	Comet                           // comet
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
	Bluejay,
	Panther,
	Cheetah,
	Lynx,
	TangorPro,
	Felix,
	Shiba,
	Husky,
	Akita,
	Tokay,
	Caiman,
	Komodo,
	Comet,
}

var deviceCodenameMap = map[Pixel]Codename{
	Pixel4:        Flame,
	Pixel4XL:      Coral,
	Pixel4a:       Sunfish,
	Pixel4a5G:     Bramble,
	Pixel5:        Redfin,
	Pixel5a:       Barbet,
	Pixel6:        Oriole,
	Pixel6Pro:     Raven,
	Pixel6a:       Bluejay,
	Pixel7:        Panther,
	Pixel7Pro:     Cheetah,
	Pixel7a:       Lynx,
	PixelTablet:   TangorPro,
	PixelFold:     Felix,
	Pixel8:        Shiba,
	Pixel8Pro:     Husky,
	Pixel8a:       Akita,
	Pixel9:        Tokay,
	Pixel9Pro:     Caiman,
	Pixel9ProXL:   Komodo,
	Pixel9ProFold: Comet,
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
	Pixel6a.String(),
	Pixel7.String(),
	Pixel7Pro.String(),
	Pixel7a.String(),
	PixelTablet.String(),
	PixelFold.String(),
	Pixel8.String(),
	Pixel8Pro.String(),
	Pixel8a.String(),
	Pixel9.String(),
	Pixel9Pro.String(),
	Pixel9ProXL.String(),
	Pixel9ProFold.String(),
	Flame.String(),
	Coral.String(),
	Sunfish.String(),
	Bramble.String(),
	Redfin.String(),
	Barbet.String(),
	Oriole.String(),
	Raven.String(),
	Bluejay.String(),
	Panther.String(),
	Cheetah.String(),
	Lynx.String(),
	TangorPro.String(),
	Felix.String(),
	Shiba.String(),
	Husky.String(),
	Akita.String(),
	Tokay.String(),
	Caiman.String(),
	Komodo.String(),
	Comet.String(),
}
