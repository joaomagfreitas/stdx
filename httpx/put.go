package httpx

import (
	"net/http"
	"net/url"
)

// Put sends an HTTP PUT request to the specified URL with optional query
// parameters and headers.
func Put(
	url string,
	body []byte,
	query url.Values,
	headers http.Header,
) (*http.Response, error) {
	req, err := Bake(url, "", http.MethodPut, body, query, headers)
	if err != nil {
		return nil, err
	}

	return http.DefaultClient.Do(req)
}
