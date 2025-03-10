package download

import (
	"io"
	"net/http"
	"os"

	"github.com/pkg/errors"
	"github.com/schollz/progressbar/v3"
)

const (
	bitsToBytesInt   int     = 1024
	bitsToBytesFloat float64 = 1024.
)

// ReadData reads data from the passed http Response body using bufSizeBytes as the
// desired buffer size.
// Returns number of bytes written to file and any error returned
func ReadData(r *http.Response, outFile string, bufSizeBytes int) (int64, error) {
	numBytes, err := readData(r.Body, outFile, r.ContentLength, bufSizeBytes)
	if err != nil {
		return numBytes, err
	}

	return numBytes, nil
}

func readData(r io.Reader, outFile string, contentLen int64, bufSizeBytes int) (int64, error) {
	out, err := os.Create(outFile)
	if err != nil {
		err = errors.WithMessagef(err, "error creating file %[1]s", outFile)
		return -1, err
	}

	var closeErr error

	defer func() {
		closeErr = out.Close()
	}()

	progressBar := progressbar.DefaultBytes(contentLen)

	buf := make([]byte, bufSizeBytes)

	numBytes, err := io.CopyBuffer(io.MultiWriter(out, progressBar), r, buf)
	if err != nil {
		err = errors.WithMessage(err, "error reading data from source")
		return -1, err
	}

	return numBytes, closeErr
}

func BytesFromMb(numBytes int) int {
	return numBytes * bitsToBytesInt * bitsToBytesInt
}

func GbFromBytes(numBytes int64) float64 {
	return float64(numBytes) / bitsToBytesFloat / bitsToBytesFloat / bitsToBytesFloat
}
