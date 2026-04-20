package routerx

import "net/http"

// Alias for [http.HandlerFunc] that returns an error.
type ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request) error
