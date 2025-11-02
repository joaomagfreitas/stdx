package writerx

import (
	"encoding/json"
	"net/http"
)

// Json attempts to marshal value [v] as a JSON document, writing it in
// the response body. If successful, it also includes `application/json`
// content type in the header fields.
func Json(w http.ResponseWriter, v any) error {
	w.Header().Set("content-type", "application/json")

	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		w.Header().Del("content-type")
	}

	return err
}
