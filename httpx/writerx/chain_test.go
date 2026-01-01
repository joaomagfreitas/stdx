package writerx_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joaomagfreitas/stdx/httpx"
	"github.com/joaomagfreitas/stdx/httpx/routerx"
	"github.com/joaomagfreitas/stdx/httpx/writerx"
)

func TestChainInvocationOrder(t *testing.T) {
	p := "foo.bar"
	doc := "<html><head><body></body></html"

	r := routerx.
		New().
		Get(p, func(w http.ResponseWriter, r *http.Request) {
			err := writerx.
				Chain(w).
				Html(doc).
				Status(300).
				Header("foo", "bar").
				Write()

			if err != nil {
				t.Fatal(err)
			}
		}).
		Build()

	s := httptest.NewServer(r)
	c := httpx.New(s.URL)

	resp, err := c.Get(p, nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	var rv string
	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	rv = string(bs)
	if doc != rv {
		t.Fail()
	}

	if resp.StatusCode != 300 {
		t.Fatal(resp.StatusCode)
	}

	if resp.Header.Get("foo") != "bar" {
		t.Fatal(resp.Header)
	}
}
