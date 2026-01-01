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

func TestHtml(t *testing.T) {
	p := "foo.bar"
	doc := "<html><head><body></body></html"

	r := routerx.
		New().
		Get(p, func(w http.ResponseWriter, r *http.Request) {
			err := writerx.Html(w, doc)
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

	ct := resp.Header.Get("content-type")
	if ct != "text/html" {
		t.Fatal(ct)
	}

	if doc != rv {
		t.Fail()
	}
}

func TestHtmlMarshalError(t *testing.T) {
	p := "foo.bar"
	doc := "<html><head><body></body></html"

	r := routerx.
		New().
		Get(p, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
			hj, ok := w.(http.Hijacker)

			if !ok {
				t.Fatal("can't hijack connection")
			}

			conn, _, err := hj.Hijack()
			if err != nil {
				t.Fatal(err)
			}

			if err = conn.Close(); err != nil {
				t.Fatal(err)
			}

			err = writerx.Html(w, doc)
			if err == nil {
				t.Fail()
			}
		}).
		Build()

	s := httptest.NewServer(r)
	c := httpx.New(s.URL)

	resp, err := c.Get(p, nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	ct := resp.Header.Get("content-type")
	if ct == "text/html" {
		t.Fatal(ct)
	}
}
