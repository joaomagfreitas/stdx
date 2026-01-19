package writerx

import (
	"net/http"
)

type chain struct {
	w  http.ResponseWriter
	wr func(w http.ResponseWriter) error
	h  map[string]string
	sc int
}

// Sets the response status code.
func (c *chain) Status(sc int) *chain {
	c.sc = sc
	return c
}

// Sets an header on the response body.
func (c *chain) Header(key, value string) *chain {
	c.h[key] = value
	return c
}

// Encodes a value in JSON on the response body.
func (c *chain) Json(v any) *chain {
	c.wr = func(w http.ResponseWriter) error {
		return Json(w, v)
	}

	c.h["content-type"] = "application/json"

	return c
}

// Encodes a value in XML on the response body.
func (c *chain) Xml(v any) *chain {
	c.wr = func(w http.ResponseWriter) error {
		return Xml(w, v)
	}

	c.h["content-type"] = "application/xml"

	return c
}

// Writes a plain text doc on the response body.
func (c *chain) Text(v any) *chain {
	c.wr = func(w http.ResponseWriter) error {
		return Text(w, v)
	}

	c.h["content-type"] = "text/plain"

	return c
}

// Writes an HTML doc on the response body.
func (c *chain) Html(doc string) *chain {
	c.wr = func(w http.ResponseWriter) error {
		return Html(w, doc)
	}

	c.h["content-type"] = "text/html"

	return c
}

// Breaks the chain by writing everything on the response body.
func (c *chain) Write() error {
	w := c.w
	wr := c.wr
	sc := c.sc
	h := c.h

	for k, v := range h {
		w.Header().Set(k, v)
	}

	if sc != 0 {
		w.WriteHeader(sc)
	}

	if wr != nil {
		return wr(w)
	}

	_, err := w.Write(nil)
	return err
}

// Chain returns a new response chain bound to [w].
//
// The returned chain buffers headers, status code, and body writer calls.
// The response is sent only when Write is invoked.
//
// Typical usage:
//
//	writerx.Chain(w).
//		Status(http.StatusCreated).
//		Header("Location", "/items/42").
//		Json(item).
//		Write()
func Chain(w http.ResponseWriter) *chain {
	return &chain{w: w, h: map[string]string{}}
}
