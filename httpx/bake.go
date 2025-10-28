package httpx

import (
	"bytes"
	"net/http"
	"net/url"
)

// Bake makes an [http.Request] from all parts required to compose it.
// If the content-type header is missing, it uses [http.DetectContentType]
// to guess what is transmitted using [body] data.
func Bake(
	baseUrl,
	endpoint,
	method string,
	body []byte,
	query url.Values,
	headers http.Header,
) (*http.Request, error) {
	u, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}

	u = u.JoinPath(endpoint)

	if len(query) > 0 {
		u.RawQuery = query.Encode()
	}

	req, err := http.NewRequest(method, u.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	if len(headers) > 0 {
		req.Header = headers
	}

	if ct := req.Header.Get("content-type"); len(ct) == 0 && body != nil {
		req.Header.Set("content-type", http.DetectContentType(body))
	}

	return req, nil
}
