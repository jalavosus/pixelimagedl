package pixelimagedl

import (
	"bytes"
	"fmt"

	"github.com/k0kubun/pp"
)

func init() {
	colorScheme := pp.ColorScheme{
		FieldName: pp.Green,
		String:    pp.Blue,
	}

	pp.SetColorScheme(colorScheme)
	pp.PrintMapTypes = false
}

type PixelImage struct {
	Version      string `json:"version" yaml:"version" pp:"-"`
	BuildNumber  string `json:"build_number" yaml:"build_number" pp:"-"`
	BuildDate    string `json:"build_date" yaml:"build_date" pp:"-"`
	BuildComment string `json:"build_comment,omitempty" yaml:"build_comment,omitempty"`
	DownloadURI  string `json:"download_uri" yaml:"download_uri" pp:"-"`
	SHA256Sum    string `json:"sha256_sum" yaml:"sha256_sum" pp:"-"`
}

func (f PixelImage) PrettyPrint() {
	fmt.Println(string(bytes.Replace(
		[]byte(pp.Sprint(f)),
		[]byte("pixelimagedl.\x1b[32mPixelImage\x1b[0m"),
		[]byte(""),
		1,
	)))
}
