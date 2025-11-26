package requests

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

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
		clientID := os.Getenv("DISCORD_CLIENT_ID")
		if clientID == "" {
			return nil, errors.New("DISCORD_CLIENT_ID is unset")
		}
		token = fmt.Sprintf(`https://discord.com/oauth2/authorize?client_id=%s
&response_type=code&redirect_uri=http%3A%2F%2Flocalhost&scope=identify`, clientID)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

// func http.NewRequest(method string, url string, body io.Reader) (*http.Request, error)
