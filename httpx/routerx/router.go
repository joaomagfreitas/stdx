package routerx

import (
	"fmt"
	"net/http"
	"strings"
)

// Alias for [http.Handler] which is the backbone of an http router in Go.
type Router = http.Handler

// Identifies a route handled by the router.
type route struct {
	Handler http.HandlerFunc
	Method  string
	Pattern string
}

// Defines a fluent interface for configuring and creating a [Router].
type routerBuilder struct {
	mux         mux
	routes      []route
	middlewares []middleware
}

// Get registers a route for HTTP GET requests matching the given pattern.
func (b *routerBuilder) Get(pattern string, handler http.HandlerFunc) *routerBuilder {
	return b.addRoute(http.MethodGet, pattern, handler)
}

// Post registers a route for HTTP POST requests matching the given pattern.
func (b *routerBuilder) Post(pattern string, handler http.HandlerFunc) *routerBuilder {
	return b.addRoute(http.MethodPost, pattern, handler)
}

// Put registers a route for HTTP PUT requests matching the given pattern.
func (b *routerBuilder) Put(pattern string, handler http.HandlerFunc) *routerBuilder {
	return b.addRoute(http.MethodPut, pattern, handler)
}

// Patch registers a route for HTTP PATCH requests matching the given pattern.
func (b *routerBuilder) Patch(pattern string, handler http.HandlerFunc) *routerBuilder {
	return b.addRoute(http.MethodPatch, pattern, handler)
}

// Delete registers a route for HTTP DELETE requests matching the given pattern.
func (b *routerBuilder) Delete(pattern string, handler http.HandlerFunc) *routerBuilder {
	return b.addRoute(http.MethodDelete, pattern, handler)
}

// Method registers a route for HTTP [method] requests matching the given pattern.
func (b *routerBuilder) Method(method, pattern string, handler http.HandlerFunc) *routerBuilder {
	return b.addRoute(method, pattern, handler)
}

// Handle registers a route for any HTTP requests matching the given pattern.
func (b *routerBuilder) Handle(pattern string, handler http.HandlerFunc) *routerBuilder {
	return b.addRoute("", pattern, handler)
}

// Use registers a middleware.
func (b *routerBuilder) Use(middleware middleware) *routerBuilder {
	b.middlewares = append(b.middlewares, middleware)
	return b
}

// Mux sets the multiplexer that will be used to handle requests.
func (b *routerBuilder) Mux(mux mux) *routerBuilder {
	b.mux = mux
	return b
}

// Build builds the router ready handle requests in a [http.Server].
func (b *routerBuilder) Build() Router {
	mux := b.mux
	if mux == nil {
		mux = http.NewServeMux()
	}

	for _, r := range b.routes {
		p := r.Pattern
		if len(r.Method) > 0 {
			p = fmt.Sprintf("%s %s", r.Method, r.Pattern)
		}

		h := chain(r.Handler, b.middlewares...)
		mux.Handle(p, h)
	}

	return mux
}

func (b *routerBuilder) addRoute(method, pattern string, handler http.HandlerFunc) *routerBuilder {
	if !strings.HasPrefix(pattern, "/") {
		pattern = "/" + pattern
	}

	r := route{Method: method, Pattern: pattern, Handler: handler}
	b.routes = append(b.routes, r)
	return b
}

// New creates and returns a new [Router] builder.
//
// It provides a fluent interface for defining HTTP routes before
// finalizing the configuration with [Build]. Use the returned builder
// to register routes using methods such as Get, Post, etc.
// You can combine another router by attaching it using [Mux].
//
// Example:
//
//	r := router.New().
//		Get("/hello", func(w http.ResponseWriter, r *http.Request) {
//			fmt.Fprintln(w, "Hello, world!")
//		}).
//		Build()
//
//	http.ListenAndServe(":8080", r)
//
// The example above creates a simple router that responds to
// GET /hello with a greeting message.
func New() *routerBuilder {
	return &routerBuilder{}
}
