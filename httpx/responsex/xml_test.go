package responsex_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joaomagfreitas/stdx/httpx"
	"github.com/joaomagfreitas/stdx/httpx/responsex"
	"github.com/joaomagfreitas/stdx/httpx/routerx"
	"github.com/joaomagfreitas/stdx/httpx/writerx"
)

type fooBar struct {
	Foo string
	Baz string
}

func TestXml(t *testing.T) {
	p := "foo.bar"
	v := fooBar{Baz: "foo.bar"}

	r := routerx.
		New().
		Get(p, func(w http.ResponseWriter, r *http.Request) {
			err := writerx.Xml(w, v)
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

	rv, err := responsex.Xml[fooBar](*resp)
	if err != nil {
		t.Fatal(err)
	}

	if rv != v {
		t.Fail()
	}
}

func TestXmlDeserializeError(t *testing.T) {
	p := "foo.bar"
	v := fooBar{Baz: "foo.bar"}

	r := routerx.
		New().
		Get(p, func(w http.ResponseWriter, r *http.Request) {
			err := writerx.Xml(w, v)
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

	_, err = responsex.Xml[func()](*resp)
	if err == nil {
		t.Fail()
	}
}
