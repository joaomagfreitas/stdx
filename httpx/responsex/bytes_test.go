package responsex_test

import (
	"net/http"
	"net/http/httptest"
	"slices"
	"testing"

	"github.com/joaomagfreitas/stdx/httpx"
	"github.com/joaomagfreitas/stdx/httpx/responsex"
	"github.com/joaomagfreitas/stdx/httpx/routerx"
)

func TestBytes(t *testing.T) {
	p := "foo.bar"
	v := []byte("foo.bar")

	r := routerx.
		New().
		Get(p, func(w http.ResponseWriter, r *http.Request) {
			_, err := w.Write(v)
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

	rv, err := responsex.Bytes(*resp)
	if err != nil {
		t.Fatal(err)
	}

	if !slices.Equal(rv, v) {
		t.Fail()
	}
}
