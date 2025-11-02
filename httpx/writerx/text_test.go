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

func TestText(t *testing.T) {
	p := "foo.bar"
	v := "baz"

	r := routerx.
		New().
		Get(p, func(w http.ResponseWriter, r *http.Request) {
			err := writerx.Text(w, v)
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

	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	rv := string(bs)

	ct := resp.Header.Get("content-type")
	if ct != "text/plain" {
		t.Fatal(ct)
	}

	if rv != v {
		t.Fail()
	}
}
