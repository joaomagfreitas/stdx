package httpx

import (
	"net/http"
	"net/url"
)

// Options sends an HTTP OPTIONS request to the specified URL with optional query
// parameters and headers.
func Options(
	url string,
	query url.Values,
	headers http.Header,
) (*http.Response, error) {
	req, err := Bake(url, "", http.MethodOptions, nil, query, headers)
	if err != nil {
		return nil, err
	}

	return http.DefaultClient.Do(req)
}
