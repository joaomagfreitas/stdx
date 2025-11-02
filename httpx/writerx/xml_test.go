package writerx_test

import (
	"encoding/xml"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joaomagfreitas/stdx/httpx"
	"github.com/joaomagfreitas/stdx/httpx/routerx"
	"github.com/joaomagfreitas/stdx/httpx/writerx"
)

type fooBar struct {
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

	var rv fooBar
	err = xml.NewDecoder(resp.Body).Decode(&rv)
	if err != nil {
		t.Fatal(err)
	}

	ct := resp.Header.Get("content-type")
	if ct != "application/xml" {
		t.Fatal(ct)
	}

	if v != rv {
		t.Fail()
	}
}

func TestXmlMarshalError(t *testing.T) {
	p := "foo.bar"
	v := func() {}

	r := routerx.
		New().
		Get(p, func(w http.ResponseWriter, r *http.Request) {
			err := writerx.Xml(w, v)
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
	if ct == "application/xml" {
		t.Fail()
	}
}
