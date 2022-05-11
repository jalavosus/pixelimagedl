package pixelimagedl

import (
	"os"
	"strconv"
)

func last[T any](data []T) T {
	return data[len(data)-1]
}

func dlBufSize() int {
	var bufSize = defaultBufSizeBytes

	if val, ok := os.LookupEnv(bufSizeEnv); ok {
		s, _ := strconv.ParseInt(val, 10, 64)
		bufSize = int(s * 1024 * 1024)
	}

	return bufSize
}

func parseInt64(s string) int64 {
	n, _ := strconv.ParseInt(s, 10, 64)
	return n
}
