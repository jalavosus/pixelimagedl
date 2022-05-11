package pixelimagedl

import (
	"encoding/json"
	"fmt"
)

type PixelImage struct {
	Version      string `json:"version" yaml:"version"`
	BuildNumber  string `json:"build_number" yaml:"build_number"`
	BuildDate    string `json:"build_date" yaml:"build_date"`
	BuildComment string `json:"build_comment" yaml:"build_comment"`
	DownloadURI  string `json:"download_uri" yaml:"download_uri"`
	SHA256Sum    string `json:"sha256_sum" yaml:"sha256_sum"`
}

func (f PixelImage) PrettyPrint() {
	d, _ := json.MarshalIndent(f, "", "  ")
	fmt.Println(string(d))
}
