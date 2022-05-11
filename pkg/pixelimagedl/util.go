package pixelimagedl

import (
	"os"

	"github.com/jalavosus/pixelimagedl/internal"
)

func dlBufSize() int {
	var bufSize = defaultBufSizeBytes

	if val, ok := os.LookupEnv(bufSizeEnv); ok {
		s := internal.ParseInt64(val)
		bufSize = int(s * 1024 * 1024)
	}

	return bufSize
}
