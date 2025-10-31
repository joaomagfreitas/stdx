package httpx

import (
	"net/http"
	"net/url"
)

// Head sends an HTTP HEAD request to the specified URL with optional query
// parameters and headers.
func Head(
	url string,
	query url.Values,
	headers http.Header,
) (*http.Response, error) {
	req, err := Bake(url, "", http.MethodHead, nil, query, headers)
	if err != nil {
		return nil, err
	}

	return http.DefaultClient.Do(req)
}
