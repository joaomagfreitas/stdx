package requestx_test

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/joaomagfreitas/stdx/httpx"
	"github.com/joaomagfreitas/stdx/httpx/requestx"
)

func TestQuery(t *testing.T) {
	k := "foo"
	v := "bar"
	q := url.Values{k: []string{v}}

	req, err := httpx.Bake("https://foo", "bar", http.MethodPost, nil, q, nil)
	if err != nil {
		t.Fatal(err)
	}

	qv := requestx.Query(*req, k)
	if qv != v {
		t.Fail()
	}
}

func TestQueryNoKey(t *testing.T) {
	q := url.Values{}
	req, err := httpx.Bake("https://foo", "bar", http.MethodPost, nil, q, nil)
	if err != nil {
		t.Fatal(err)
	}

	v := requestx.Query(*req, "foo")
	if v != "" {
		t.Fail()
	}

}
