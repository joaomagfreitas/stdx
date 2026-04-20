package routerx

import "net/http"

// An [http.HandlerFunc] that returns an error.
type ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request) error

// A middleware func that accepts an [ErrorHandlerFunc] in return for a [http.HandlerFunc].
type ErrorHandler func(handlerFunc ErrorHandlerFunc) http.HandlerFunc
