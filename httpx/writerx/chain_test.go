package writerx_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joaomagfreitas/stdx/httpx"
	"github.com/joaomagfreitas/stdx/httpx/routerx"
	"github.com/joaomagfreitas/stdx/httpx/writerx"
)

func TestChainInvocationOrder(t *testing.T) {
	p := "foo.bar"
	doc := fooBar{Baz: "bar"}

	r := routerx.
		New().
		Get(p, func(w http.ResponseWriter, r *http.Request) {
			err := writerx.
				Chain(w).
				Text("foo bar").
				Html("<foo>bar</bar>").
				Xml("<foo>bar</bar>").
				Json(doc).
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

	var rdoc fooBar
	err = json.NewDecoder(resp.Body).Decode(&rdoc)
	if err != nil {
		t.Fatal(err)
	}

	if doc != rdoc {
		t.Fatal(rdoc)
	}

	if resp.StatusCode != 300 {
		t.Fatal(resp.StatusCode)
	}

	if resp.Header.Get("foo") != "bar" {
		t.Fatal(resp.Header)
	}

	if resp.Header.Get("content-type") != "application/json" {
		t.Fatal(resp.Header)
	}
}
