package responsex

import (
	"io"
	"net/http"
)

// Bytes reads response body as binary data.
func Bytes(r http.Response) ([]byte, error) {
	return io.ReadAll(r.Body)
}
