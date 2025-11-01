package requestx

import (
	"io"
	"net/http"
)

// Bytes reads request body as binary data.
func Bytes(r http.Request) ([]byte, error) {
	return io.ReadAll(r.Body)
}
