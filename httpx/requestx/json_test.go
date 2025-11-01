package requestx_test

import (
	"encoding/json"
	"net/http"
	"slices"
	"testing"

	"github.com/joaomagfreitas/stdx/httpx"
	"github.com/joaomagfreitas/stdx/httpx/requestx"
)

func TestJson(t *testing.T) {
	s := []string{"foo", "bar"}
	b, err := json.Marshal(s)
	if err != nil {
		t.Fatal(err)
	}

	req, err := httpx.Bake("https://foo", "bar", http.MethodPost, b, nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	v, err := requestx.Json[[]string](*req)
	if err != nil {
		t.Fatal(err)
	}

	if !slices.Equal(s, v) {
		t.Fail()
	}
}

func TestJsonParseError(t *testing.T) {
	b := []byte("not a json")
	req, err := httpx.Bake("https://foo", "bar", http.MethodPost, b, nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	_, err = requestx.Json[string](*req)
	if err == nil {
		t.Fail()
	}
}

func TestJsonDeserializeError(t *testing.T) {
	s := []string{"foo", "bar"}
	b, err := json.Marshal(s)
	if err != nil {
		t.Fatal(err)
	}

	req, err := httpx.Bake("https://foo", "bar", http.MethodPost, b, nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	_, err = requestx.Json[[]int](*req)
	if err == nil {
		t.Fail()
	}
}
