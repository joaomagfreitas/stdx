package writerx

import (
	"net/http"
)

// Html attempts to write value [doc] in the response body, as an HTML document.
// If successful, it also includes `text/html` content type in the header fields.
func Html(w http.ResponseWriter, doc string) error {
	w.Header().Set("content-type", "text/html")
	_, err := w.Write([]byte(doc))

	if err != nil {
		w.Header().Del("content-type")
	}

	return err
}
