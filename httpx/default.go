package httpx

import (
	"net/http"
	"net/url"
)

type defaultClient struct {
	client  *http.Client
	baseUrl string
}

func (c defaultClient) Get(endpoint string, query url.Values, headers http.Header) (*http.Response, error) {
	req, err := Bake(c.baseUrl, endpoint, http.MethodGet, nil, query, headers)
	if err != nil {
		return nil, err
	}

	return c.client.Do(req)
}

func (c defaultClient) Post(endpoint string, body []byte, query url.Values, headers http.Header) (*http.Response, error) {
	req, err := Bake(c.baseUrl, endpoint, http.MethodPost, body, query, headers)
	if err != nil {
		return nil, err
	}

	return c.client.Do(req)
}

func (c defaultClient) Put(endpoint string, body []byte, query url.Values, headers http.Header) (*http.Response, error) {
	req, err := Bake(c.baseUrl, endpoint, http.MethodPut, body, query, headers)
	if err != nil {
		return nil, err
	}

	return c.client.Do(req)
}

func (c defaultClient) Patch(endpoint string, body []byte, query url.Values, headers http.Header) (*http.Response, error) {
	req, err := Bake(c.baseUrl, endpoint, http.MethodPatch, body, query, headers)
	if err != nil {
		return nil, err
	}

	return c.client.Do(req)
}

func (c defaultClient) Delete(endpoint string, query url.Values, headers http.Header) (*http.Response, error) {
	req, err := Bake(c.baseUrl, endpoint, http.MethodDelete, nil, query, headers)
	if err != nil {
		return nil, err
	}

	return c.Do(req)
}

func (c defaultClient) Head(endpoint string, query url.Values, headers http.Header) (*http.Response, error) {
	req, err := Bake(c.baseUrl, endpoint, http.MethodHead, nil, query, headers)
	if err != nil {
		return nil, err
	}

	return c.Do(req)
}

func (c defaultClient) Options(endpoint string, query url.Values, headers http.Header) (*http.Response, error) {
	req, err := Bake(c.baseUrl, endpoint, http.MethodOptions, nil, query, headers)
	if err != nil {
		return nil, err
	}

	return c.Do(req)
}

func (c defaultClient) Do(request *http.Request) (*http.Response, error) {
	return c.client.Do(request)
}

// Builds a [Client] that can be used to send HTTP requests towards a server
// that listens for requests in [baseUrl]. Uses Go [http.DefaultClient] for all
// operations.
func New(
	baseUrl string,
) Client {
	c := defaultClient{
		client:  http.DefaultClient,
		baseUrl: baseUrl,
	}

	return &c
}
