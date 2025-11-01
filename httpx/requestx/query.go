package requestx

import "net/http"

// Query extracts the value of the query parameter associated to [key].
func Query(r http.Request, key string) string {
	return r.URL.Query().Get(key)
}
