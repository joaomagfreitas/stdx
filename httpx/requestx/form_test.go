package requestx_test

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/joaomagfreitas/stdx/httpx"
	"github.com/joaomagfreitas/stdx/httpx/requestx"
)

func TestFormUrl(t *testing.T) {
	k := "foo"
	v := "bar"
	q := url.Values{k: []string{v}}

	req, err := httpx.Bake("https://foo", "bar", http.MethodPost, nil, q, nil)
	if err != nil {
		t.Fatal(err)
	}

	qv := requestx.Form(*req, k)
	if qv != v {
		t.Fail()
	}
}

func TestFormBodyUrlEncoded(t *testing.T) {
	k := "foo"
	v := "bar"
	b := fmt.Appendf(nil, "%s=%s", k, v)
	h := http.Header{http.CanonicalHeaderKey("content-type"): []string{"application/x-www-form-urlencoded "}}

	req, err := httpx.Bake("https://foo", "bar", http.MethodPost, b, nil, h)
	if err != nil {
		t.Fatal(err)
	}

	qv := requestx.Form(*req, k)
	if qv != v {
		t.Fail()
	}
}
