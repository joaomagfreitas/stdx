package routerx

import "net/http"

// Abstraction over [http.ServerMux] so you can pass different
// implementations of [http.Handler] to [routerBuilder].
type mux interface {
	http.Handler
	Handle(pattern string, handler http.Handler)
}
