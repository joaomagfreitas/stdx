package httpx

import (
	"net/http"
	"net/url"
)

// Get sends an HTTP GET request to the specified URL with optional query
// parameters and headers.
func Get(
	url string,
	query url.Values,
	headers http.Header,
) (*http.Response, error) {
	req, err := Bake(url, "", http.MethodGet, nil, query, headers)
	if err != nil {
		return nil, err
	}

	return http.DefaultClient.Do(req)
}
