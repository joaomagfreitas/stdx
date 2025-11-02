package writerx

import (
	"fmt"
	"net/http"
)

// Text attempts to marshal value [v] as a text document, writing it in
// the response body. If successful, it also includes `text/plain`
// content type in the header fields.
func Text(w http.ResponseWriter, v any) error {
	w.Header().Set("content-type", "text/plain")

	_, err := fmt.Fprintf(w, "%s", v)
	if err != nil {
		w.Header().Del("content-type")
	}

	return nil
}
