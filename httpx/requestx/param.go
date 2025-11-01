package requestx

import "net/http"

// Param extracts the value of the path parameter associated to [name]
// that was matched by a [net/http] router.
func Param(r http.Request, name string) string {
	return r.PathValue(name)
}
