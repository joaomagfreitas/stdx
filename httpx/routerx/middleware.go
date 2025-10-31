package routerx

import "net/http"

// Alias for a middleware func that can be combined in a [http.Handler] router.
type middleware = func(next http.Handler) http.Handler

// Chains multiple outbound middlewares to an [http.Handler].
func chain(handler http.Handler, middlewares ...middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}

	return handler
}
