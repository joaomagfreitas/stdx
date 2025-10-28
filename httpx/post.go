package httpx

import (
	"net/http"
	"net/url"
)

// Post sends an HTTP POST request to the specified URL with optional query
// parameters and headers.
func Post(
	url string,
	body []byte,
	query url.Values,
	headers http.Header,
) (*http.Response, error) {
	req, err := Bake(url, "", http.MethodPost, body, query, headers)
	if err != nil {
		return nil, err
	}

	return http.DefaultClient.Do(req)
}
