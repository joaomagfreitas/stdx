package responsex

import (
	"encoding/xml"
	"net/http"
)

// Xml reads response body as a XML document and
// deserializes as a value of type [T].
func Xml[T any](r http.Response) (T, error) {
	var v T
	return v, xml.NewDecoder(r.Body).Decode(&v)
}
