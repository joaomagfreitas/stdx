package requestx

import "net/http"

// Form extracts the value of the request form associated to [key].
func Form(r *http.Request, key string) string {
	return r.FormValue(key)
}
