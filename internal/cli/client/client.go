package client

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/oauth"
)

func SendRequestAndReadResponse(url *url.URL, enableAuth bool, method string, body io.Reader) ([]byte, error) {
	client := http.Client{}
	var req *http.Request
	var err error
	if enableAuth {
		req, err = oauth.NewRequestWithAuth(method, url.String(), body)
	} else {
		req, err = http.NewRequest(method, url.String(), body)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to construct request: %w", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP %s", res.Status)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	return data, nil
}
