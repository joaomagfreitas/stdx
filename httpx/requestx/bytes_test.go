package requestx_test

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/joaomagfreitas/stdx/httpx"
	"github.com/joaomagfreitas/stdx/httpx/requestx"
)

func TestBytes(t *testing.T) {
	b := []byte("foo.bar")
	req, err := httpx.Bake("https://foo", "bar", http.MethodPost, b, nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	bs, err := requestx.Bytes(*req)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(bs, b) {
		t.Fail()
	}
}
