package writerx

import (
	"encoding/xml"
	"net/http"
)

// Xml attempts to marshal value [v] as a XML document, writing it in
// the response body. If successful, it also includes `application/xml`
// content type in the header fields.
func Xml(w http.ResponseWriter, v any) error {
	w.Header().Set("content-type", "application/xml")

	err := xml.NewEncoder(w).Encode(v)
	if err != nil {
		w.Header().Del("content-type")
	}

	return err
}
