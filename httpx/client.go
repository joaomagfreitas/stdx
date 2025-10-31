package httpx

import (
	"net/http"
	"net/url"
)

// Client abstracts HTTP methods for standard API operations. Each method accepts
// an endpoint string instead of an URL, so callers don't need to pass the full URL
// for each method.
type Client interface {
	// Get sends an HTTP GET request.
	Get(endpoint string, query url.Values, headers http.Header) (*http.Response, error)

	// Post sends an HTTP POST request with the given body. If `Content-Type`
	// header is not passed in [headers] value, it uses [http.DetectContentType]
	Post(endpoint string, body []byte, query url.Values, headers http.Header) (*http.Response, error)

	// Put sends an HTTP PUT request with the given body.
	Put(endpoint string, body []byte, query url.Values, headers http.Header) (*http.Response, error)

	// Patch sends an HTTP PATCH request.
	Patch(endpoint string, body []byte, query url.Values, headers http.Header) (*http.Response, error)

	// Delete sends an HTTP DELETE request.
	Delete(endpoint string, query url.Values, headers http.Header) (*http.Response, error)

	// Delete sends an HTTP HEAD request.
	Head(endpoint string, query url.Values, headers http.Header) (*http.Response, error)

	// Delete sends an HTTP OPTIONS request.
	Options(endpoint string, query url.Values, headers http.Header) (*http.Response, error)

	// Do executes an already composed HTTP request.
	Do(request *http.Request) (*http.Response, error)
}
