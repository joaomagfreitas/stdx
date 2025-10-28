package httpx_test

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/joaomagfreitas/stdx/httpx"
)

func TestBakeUrlComposition(t *testing.T) {
	testCases := []struct {
		desc     string
		baseUrl  string
		endpoint string
		query    url.Values
		composed string
	}{
		{
			desc:     "joins endpoint to base url",
			baseUrl:  "https://foo",
			endpoint: "bar",
			composed: "https://foo/bar",
		},
		{
			desc:     "strips leading slash of endpoint",
			baseUrl:  "https://foo",
			endpoint: "/bar",
			composed: "https://foo/bar",
		},
		{
			desc:     "strips trailing slash of endpoint",
			baseUrl:  "https://foo/",
			endpoint: "bar",
			composed: "https://foo/bar",
		},
		{
			desc:     "appends query parameters",
			baseUrl:  "https://foo",
			endpoint: "bar",
			composed: "https://foo/bar?baz=true&c=123",
			query: url.Values{
				"baz": []string{"true"},
				"c":   []string{"123"},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			req, err := httpx.Bake(tC.baseUrl, tC.endpoint, http.MethodGet, nil, tC.query, nil)
			if err != nil {
				t.Fatal(err)
			}

			if req.URL.String() != tC.composed {
				t.Fail()
			}
		})
	}
}

func TestBakeContentType(t *testing.T) {
	testCases := []struct {
		headers     http.Header
		desc        string
		contentType string
		body        []byte
	}{
		{
			desc:        "does not apply content-type header if body is nil",
			body:        nil,
			contentType: "",
		},
		{
			desc:        "auto applies content-type header for body content",
			body:        []byte("foo.bar"),
			contentType: "text/plain; charset=utf-8",
		},
		{
			desc:        "auto applies binary content-type header if body content is unrecognized",
			body:        []byte{0, 1, 2},
			contentType: "application/octet-stream",
		},
		{
			desc:        "does not auto-apply content-type header if already present in headers",
			body:        nil,
			contentType: "foo.bar",
			headers:     http.Header{http.CanonicalHeaderKey("content-type"): []string{"foo.bar"}},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			req, err := httpx.Bake("", "", "", tC.body, nil, tC.headers)
			if err != nil {
				t.Fatal(err)
			}

			ct := req.Header.Get("content-type")
			if ct != tC.contentType {
				t.Fail()
			}
		})
	}
}

func TestBakeErrors(t *testing.T) {
	testCases := []struct {
		desc     string
		baseUrl  string
		endpoint string
		method   string
	}{
		{
			desc:     "fails baking url if missing schema",
			baseUrl:  "://foo",
			endpoint: "bar",
		},
		{
			desc:     "fails baking request if using invalid method",
			baseUrl:  "https://foo",
			endpoint: "bar",
			method:   "0x\f",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			_, err := httpx.Bake(tC.baseUrl, tC.endpoint, tC.method, nil, nil, nil)
			if err == nil {
				t.Fail()
			}
		})
	}
}
