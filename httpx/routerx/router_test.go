package routerx_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/joaomagfreitas/stdx/httpx"
	"github.com/joaomagfreitas/stdx/httpx/routerx"
)

func TestRouterMethodHandlers(t *testing.T) {
	m := "foo"
	p := "bar"
	h := func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(200) }
	r := routerx.
		New().
		Get(p, h).
		Post(p, h).
		Put(p, h).
		Patch(p, h).
		Delete(p, h).
		Method(m, p, h).
		Build()

	s := httptest.NewServer(r)
	c := httpx.New(s.URL)

	testCases := []struct {
		desc   string
		method string
	}{
		{
			desc:   "handles GET method",
			method: http.MethodGet,
		},
		{
			desc:   "handles POST method",
			method: http.MethodPost,
		},
		{
			desc:   "handles PUT method",
			method: http.MethodPut,
		},
		{
			desc:   "handles PATCH method",
			method: http.MethodPatch,
		},
		{
			desc:   "handles DELETE method",
			method: http.MethodDelete,
		},
		{
			desc:   "handles other methods",
			method: m,
		},
		{
			desc:   "handles other methods",
			method: m,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			req, err := httpx.Bake(s.URL, p, tC.method, nil, nil, nil)
			if err != nil {
				t.Fatal(err)
			}

			resp, err := c.Do(req)
			if err != nil {
				t.Fatal(err)
			}

			if resp.StatusCode != 200 {
				t.Fail()
			}
		})
	}
}

func TestRouterMethodCaseSensitive(t *testing.T) {
	m := "foo"
	p := "bar"
	h := func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(200) }
	r := routerx.
		New().
		Method(m, p, h).
		Build()

	s := httptest.NewServer(r)
	c := httpx.New(s.URL)

	req, err := httpx.Bake(s.URL, p, m, nil, nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Fail()
	}

	req, err = httpx.Bake(s.URL, p, strings.ToUpper(m), nil, nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err = c.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode == 200 {
		t.Fail()
	}
}

func TestRouterHandleAnyRequest(t *testing.T) {
	p := "bar"
	h := func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(200) }
	r := routerx.
		New().
		Handle(p, h).
		Build()

	s := httptest.NewServer(r)
	c := httpx.New(s.URL)

	ms := []string{
		"foo",
		"bar",
		"baz",
		"Foo.Bar.Baz",
		"get",
		"GET",
	}

	for _, m := range ms {
		req, err := httpx.Bake(s.URL, p, m, nil, nil, nil)
		if err != nil {
			t.Fatal(err)
		}

		resp, err := c.Do(req)
		if err != nil {
			t.Fatal(err)
		}

		if resp.StatusCode != 200 {
			t.Fail()
		}
	}
}

func TestRouterMiddleware(t *testing.T) {
	m := "foo"
	p := "bar"
	ct := "baz"
	h := func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(200) }
	mw := func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.Header().Set("content-type", ct) }),
		)
	}
	r := routerx.
		New().
		Method(m, p, h).
		Use(mw).
		Build()

	s := httptest.NewServer(r)
	c := httpx.New(s.URL)

	req, err := httpx.Bake(s.URL, p, m, nil, nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Fatal(resp)
	}

	if resp.Header.Get("content-type") != ct {
		t.Fail()
	}
}

func TestRouterMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.Handle("/foo.bar", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}))

	m := "foo"
	p := "bar"
	h := func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(200) }
	r := routerx.
		New().
		Method(m, p, h).
		Mux(mux).
		Build()

	s := httptest.NewServer(r)
	c := httpx.New(s.URL)

	ps := []string{p, "/foo.bar"}

	for _, p := range ps {
		req, err := httpx.Bake(s.URL, p, m, nil, nil, nil)
		if err != nil {
			t.Fatal(err)
		}

		resp, err := c.Do(req)
		if err != nil {
			t.Fatal(err)
		}

		if resp.StatusCode != 200 {
			t.Fatal(resp)
		}
	}

}
