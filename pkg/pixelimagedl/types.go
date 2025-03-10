package pixelimagedl

import (
	"bytes"
	"fmt"

	"github.com/k0kubun/pp/v3"
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
	Version      string `json:"version" yaml:"version" pp:"version"`
	BuildNumber  string `json:"build_number" yaml:"build_number" pp:"build_number"`
	BuildDate    string `json:"build_date" yaml:"build_date" pp:"build_date"`
	BuildComment string `json:"build_comment,omitempty" yaml:"build_comment,omitempty" pp:"build_comment,omitempty"`
	DownloadURI  string `json:"download_uri" yaml:"download_uri" pp:"download_uri"`
	SHA256Sum    string `json:"sha256_sum" yaml:"sha256_sum" pp:"sha256_sum"`
}

func (f *PixelImage) PrettyPrint() {
	out := new(bytes.Buffer)
	_, _ = pp.Fprint(out, f)

	fmt.Println(string(bytes.Replace(
		out.Bytes(),
		[]byte("pixelimagedl.\x1b[32mPixelImage\x1b[0m"),
		[]byte(""),
		1,
	)))
}
