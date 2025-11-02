package responsex

import (
	"encoding/json"
	"net/http"
)

// Json reads response body as a JSON document and
// deserializes as a value of type [T].
func Json[T any](r http.Response) (T, error) {
	var v T
	return v, json.NewDecoder(r.Body).Decode(&v)
}
