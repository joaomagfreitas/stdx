package requestx_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joaomagfreitas/stdx/httpx"
	"github.com/joaomagfreitas/stdx/httpx/requestx"
	"github.com/joaomagfreitas/stdx/httpx/routerx"
)

func TestParam(t *testing.T) {
	p := "/{id}"
	n := "id"
	ep := "foo"

	r := routerx.
		New().
		Get(p, func(w http.ResponseWriter, r *http.Request) {
			v := requestx.Param(*r, n)
			if v != ep {
				t.Fail()
			}
		}).
		Build()

	s := httptest.NewServer(r)
	c := httpx.New(s.URL)

	_, err := c.Get(ep, nil, nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestParamNotFound(t *testing.T) {
	p := "/{id2}"
	n := "id"
	ep := "foo"

	r := routerx.
		New().
		Get(p, func(w http.ResponseWriter, r *http.Request) {
			v := requestx.Param(*r, n)
			if v == ep {
				t.Fail()
			}
		}).
		Build()

	s := httptest.NewServer(r)
	c := httpx.New(s.URL)

	_, err := c.Get(ep, nil, nil)
	if err != nil {
		t.Fatal(err)
	}

}
