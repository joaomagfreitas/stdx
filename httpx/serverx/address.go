package serverx

import (
	"net"
	"os"
	"strings"

	"github.com/joaomagfreitas/stdx/slicesx"
)

// common env variables for the http server host.
var hosts = []string{
	"host",
	"http.host",
	"http_host",
	"server.host",
	"server_host",
}

// common env variables for the http server port.
var ports = []string{
	"port",
	"http.port",
	"http_port",
	"server.port",
	"server_port",
}

var defaultHost = "::"
var defaultPort = "8080"

// Address resolves a ready to serve http address, using common env variables.
// Defaults to "::" and "8080" (host, port) if the environment variables are not defined.
func Address() string {
	hs := append(hosts, slicesx.Map(hosts, func(h string) string { return strings.ToUpper(h) })...)
	ps := append(ports, slicesx.Map(ports, func(p string) string { return strings.ToUpper(p) })...)

	h, ok := slicesx.First(hs, func(h string) bool { return len(os.Getenv(h)) > 0 })
	if ok {
		h = os.Getenv(h)
	} else {
		h = defaultHost
	}

	p, ok := slicesx.First(ps, func(p string) bool { return len(os.Getenv(p)) > 0 })
	if ok {
		p = os.Getenv(p)
	} else {
		p = defaultPort
	}

	return net.JoinHostPort(h, p)
}
