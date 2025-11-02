package requestx_test

import (
	"encoding/xml"
	"net/http"
	"testing"

	"github.com/joaomagfreitas/stdx/httpx"
	"github.com/joaomagfreitas/stdx/httpx/requestx"
)

type fooBar struct {
	Foo string
	Bar string
}

func TestXml(t *testing.T) {
	s := fooBar{Foo: "bar"}
	b, err := xml.Marshal(s)
	if err != nil {
		t.Fatal(err)
	}

	req, err := httpx.Bake("https://foo", "bar", http.MethodPost, b, nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	v, err := requestx.Xml[fooBar](*req)
	if err != nil {
		t.Fatal(err)
	}

	if v != s {
		t.Fail()
	}
}

func TestXmlParseError(t *testing.T) {
	b := []byte("not a xml doc")
	req, err := httpx.Bake("https://foo", "bar", http.MethodPost, b, nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	_, err = requestx.Xml[fooBar](*req)
	if err == nil {
		t.Fail()
	}
}

func TestXmlDeserializeError(t *testing.T) {
	s := fooBar{Foo: "bar"}
	b, err := xml.Marshal(s)
	if err != nil {
		t.Fatal(err)
	}

	req, err := httpx.Bake("https://foo", "bar", http.MethodPost, b, nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	_, err = requestx.Xml[func()](*req)
	if err == nil {
		t.Fail()
	}
}
