package serverx_test

import (
	"os"
	"testing"

	"github.com/joaomagfreitas/stdx/httpx/serverx"
)

func TestAddressEnvVariables(t *testing.T) {
	defer os.Clearenv()

	if err := os.Setenv("http.host", "foo"); err != nil {
		t.Fatal(err)
	}

	if err := os.Setenv("http.port", "bar"); err != nil {
		t.Fatal(err)
	}

	addr := serverx.Address()
	if addr != "foo:bar" {
		t.Fail()
	}
}

func TestAddressFallback(t *testing.T) {
	testCases := []struct {
		desc string
		host []string
		port []string
		addr string
	}{
		{
			desc: "uses fallback if host env key is missing",
			host: nil,
			port: []string{"port", "bar"},
			addr: "[::]:bar",
		},
		{
			desc: "uses fallback if port env key is missing",
			host: []string{"host", "foo"},
			port: nil,
			addr: "foo:8080",
		},
		{
			desc: "uses fallback if both host and port env key are missing",
			host: nil,
			port: nil,
			addr: "[::]:8080",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			defer os.Clearenv()

			if len(tC.host) == 2 {
				if err := os.Setenv(tC.host[0], tC.host[1]); err != nil {
					t.Fatal(err)
				}
			}

			if len(tC.port) == 2 {
				if err := os.Setenv(tC.port[0], tC.port[1]); err != nil {
					t.Fatal(err)
				}
			}

			addr := serverx.Address()
			if addr != tC.addr {
				t.Fail()
			}
		})
	}
}
