package responsex_test

import (
	"net/http"
	"net/http/httptest"
	"slices"
	"testing"

	"github.com/joaomagfreitas/stdx/httpx"
	"github.com/joaomagfreitas/stdx/httpx/responsex"
	"github.com/joaomagfreitas/stdx/httpx/routerx"
	"github.com/joaomagfreitas/stdx/httpx/writerx"
)

func TestJson(t *testing.T) {
	p := "foo.bar"
	v := []string{"foo", "bar"}

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

	rv, err := responsex.Json[[]string](*resp)
	if err != nil {
		t.Fatal(err)
	}

	if !slices.Equal(rv, v) {
		t.Fail()
	}
}

func TestJsonDeserializeError(t *testing.T) {
	p := "foo.bar"
	v := []string{"foo", "bar"}

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

	_, err = responsex.Json[[]int](*resp)
	if err == nil {
		t.Fatal(err)
	}
}
