package client

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func SendRequestAndReadResponse(url *url.URL, method string, body io.Reader) ([]byte, error) {
	client := http.Client{}
	req, err := http.NewRequest(method, url.String(), body)
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
