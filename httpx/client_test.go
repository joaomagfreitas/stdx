package httpx_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/joaomagfreitas/stdx/httpx"
)

func TestSmokePingPong(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bs, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}

		w.Header().Set("x-body", string(bs))
	}))

	c := httpx.New(s.URL)

	testCases := []struct {
		reqFn func() (*http.Response, error)
		desc  string
		body  string
	}{
		{
			desc: "get",
			reqFn: func() (*http.Response, error) {
				return httpx.Get(s.URL, nil, nil)
			},
		},
		{
			desc: "post",
			body: "foo.bar",
			reqFn: func() (*http.Response, error) {
				return httpx.Post(s.URL, []byte("foo.bar"), nil, nil)
			},
		},
		{
			desc: "put",
			body: "[0,1,2]",
			reqFn: func() (*http.Response, error) {
				b := mustEncodeJson(t, []int{0, 1, 2})
				h := http.Header{http.CanonicalHeaderKey("content-type"): []string{"application/json"}}
				return httpx.Put(s.URL, b, nil, h)
			},
		},
		{
			desc: "patch",
			reqFn: func() (*http.Response, error) {
				q := url.Values{"foo": []string{"bar"}}
				return httpx.Patch(s.URL, nil, q, nil)
			},
		},
		{
			desc: "delete",
			reqFn: func() (*http.Response, error) {
				return httpx.Delete(s.URL, nil, nil)
			},
		},

		{
			desc: "get (client)",
			reqFn: func() (*http.Response, error) {
				return c.Get(s.URL, nil, nil)
			},
		},
		{
			desc: "post (client)",
			body: "foo.bar",
			reqFn: func() (*http.Response, error) {
				return c.Post(s.URL, []byte("foo.bar"), nil, nil)
			},
		},
		{
			desc: "put (client)",
			body: "[0,1,2]",
			reqFn: func() (*http.Response, error) {
				b := mustEncodeJson(t, []int{0, 1, 2})
				h := http.Header{http.CanonicalHeaderKey("content-type"): []string{"application/json"}}
				return c.Put(s.URL, b, nil, h)
			},
		},
		{
			desc: "patch (client)",
			reqFn: func() (*http.Response, error) {
				q := url.Values{"foo": []string{"bar"}}
				return c.Patch(s.URL, nil, q, nil)
			},
		},
		{
			desc: "delete (client)",
			reqFn: func() (*http.Response, error) {
				return c.Delete(s.URL, nil, nil)
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			resp, err := tC.reqFn()
			if err != nil {
				t.Fatal(err)
			}

			b := resp.Header.Get("x-body")
			if b != tC.body {
				t.Fail()
			}
		})
	}
}

func mustEncodeJson(t *testing.T, v any) []byte {
	bs, err := json.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}

	return bs
}
