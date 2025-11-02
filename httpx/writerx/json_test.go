package writerx_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"slices"
	"testing"

	"github.com/joaomagfreitas/stdx/httpx"
	"github.com/joaomagfreitas/stdx/httpx/routerx"
	"github.com/joaomagfreitas/stdx/httpx/writerx"
)

func TestJson(t *testing.T) {
	p := "foo.bar"
	v := []string{"baz"}

	r := routerx.
		New().
		Get(p, func(w http.ResponseWriter, r *http.Request) {
			err := writerx.Json(w, v)
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

	var rv []string
	err = json.NewDecoder(resp.Body).Decode(&rv)
	if err != nil {
		t.Fatal(err)
	}

	ct := resp.Header.Get("content-type")
	if ct != "application/json" {
		t.Fatal(ct)
	}

	if !slices.Equal(rv, v) {
		t.Fail()
	}
}

func TestJsonMarshalError(t *testing.T) {
	p := "foo.bar"
	v := func() {}

	r := routerx.
		New().
		Get(p, func(w http.ResponseWriter, r *http.Request) {
			err := writerx.Json(w, v)
			if err == nil {
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

	ct := resp.Header.Get("content-type")
	if ct == "application/json" {
		t.Fail()
	}
}
