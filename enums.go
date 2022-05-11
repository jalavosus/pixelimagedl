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
	Pixel4    Pixel = iota // Pixel 4
	Pixel4XL               // Pixel 4 XL
	Pixel4a                // Pixel 4a
	Pixel4a5G              // Pixel 4a 5G
	Pixel5                 // Pixel 5
	Pixel5a                // Pixel 5a
	Pixel6                 // Pixel 6
	Pixel6Pro              // Pixel 6 Pro
	Unknown                // unknown
)

const (
	Flame   Codename = iota // flame
	Coral                   // coral
	Sunfish                 // sunfish
	Bramble                 // bramble
	Redfin                  // redfin
	Barbet                  // barbet
	Oriole                  // oriole
	Raven                   // raven
)

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
