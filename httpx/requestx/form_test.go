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
	m := map[string]string{
		"foo":   "bar",
		"baz":   "lorem",
		"ipsum": "doli",
	}

	q := url.Values{}
	for k, v := range m {
		q.Add(k, v)
	}

	req, err := httpx.Bake("https://foo", "bar", http.MethodGet, nil, q, nil)
	if err != nil {
		t.Fatal(err)
	}

	for k, v := range m {
		qv := requestx.Form(req, k)
		if qv != v {
			t.Fatal(qv)
		}
	}
}

func TestFormBodyUrlEncoded(t *testing.T) {
	m := map[string]string{
		"foo":   "bar",
		"baz":   "lorem",
		"ipsum": "doli",
	}

	var b []byte
	for k, v := range m {
		if len(b) == 0 {
			b = fmt.Appendf(nil, "%s=%s", k, v)
		} else {
			b = fmt.Appendf(b, "&%s=%s", k, v)
		}
	}

	h := http.Header{http.CanonicalHeaderKey("content-type"): []string{"application/x-www-form-urlencoded "}}

	req, err := httpx.Bake("https://foo", "bar", http.MethodPost, b, nil, h)
	if err != nil {
		t.Fatal(err)
	}

	for k, v := range m {
		qv := requestx.Form(req, k)
		if qv != v {
			t.Fatal(qv)
		}
	}
}
