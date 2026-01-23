package serverx

import (
	"net/http"

	"github.com/joaomagfreitas/stdx/logx"
)

// ListenAndServe is analogous to [http.ListenAndServe], but includes a trace log
// of the address which the http server will be listening for requestrs.
func ListenAndServe(addr string, h http.Handler) error {
	logx.Tracef("Starting http server on http://%s", addr)

	return http.ListenAndServe(addr, h)
}
