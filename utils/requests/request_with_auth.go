package requests

import (
	"fmt"
	"io"
	"net/http"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/config"
)

func NewRequestWithAuth(method, url string, body io.Reader) (*http.Request, error) {
	cfg := config.Load()
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	var token string
	if cfg.Env == "development" {
		token = "dev-token"
	} else {
		// TODO
		token = "asdf"
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

// func http.NewRequest(method string, url string, body io.Reader) (*http.Request, error)
