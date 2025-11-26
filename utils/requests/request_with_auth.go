package requests

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/cli/browser"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/config"
)

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

func NewRequestWithAuth(method, targetURL string, body io.Reader) (*http.Request, error) {
	cfg := config.Load()
	req, err := http.NewRequest(method, targetURL, body)
	if err != nil {
		return nil, err
	}

	var token string
	if cfg.Env == "development" {
		token = "dev-token"
	} else {
		// Production OAuth2 Flow
		clientID := os.Getenv("DISCORD_CLIENT_ID")
		if clientID == "" {
			return nil, errors.New("DISCORD_CLIENT_ID is unset")
		}
		clientSecret := os.Getenv("DISCORD_CLIENT_SECRET")
		if clientSecret == "" {
			return nil, errors.New("DISCORD_CLIENT_SECRET is unset")
		}
		// TODO: check that this port isn't being used first
		const redirectURI = "http://localhost:8888"
		const scope = "identify"

		tokenChan := make(chan string)
		errChan := make(chan error)
		mux := http.NewServeMux()
		server := &http.Server{Addr: ":8888", Handler: mux}
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			code := r.URL.Query().Get("code")
			if code == "" {
				fmt.Fprintln(w, "No code found")
				return
			}
			fmt.Fprintf(w, "Got code! You can close this window.")

			data := url.Values{}
			// data.Set("client_id", clientID)
			data.Set("client_secret", clientSecret)
			data.Set("grant_type", "authorization_code")
			data.Set("code", code)
			data.Set("redirect_uri", redirectURI)

			resp, err := http.PostForm("https://discord.com/api/oauth2/token", data)
			if err != nil {
				errChan <- fmt.Errorf("failed to exchange token: %w", err)
				return
			}
			defer resp.Body.Close()

			respBody, _ := io.ReadAll(resp.Body)

			if resp.StatusCode != http.StatusOK {
				errChan <- fmt.Errorf("discord API error: %s", string(respBody))
			}

			var tokenResp TokenResponse
			if err := json.Unmarshal(respBody, &tokenResp); err != nil {
				errChan <- fmt.Errorf("JSON parse error: %w", err)
			}

			tokenChan <- tokenResp.AccessToken
		})

		// So server doesn't block
		go func() {
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				errChan <- err
			}
		}()

		// Human needs to open browser to get the token
		// A web client can do this more gracefully, but we have to do this since we're on a CLI
		params := url.Values{}
		params.Add("client_id", clientID)
		params.Add("redirect_uri", redirectURI)
		params.Add("scope", scope)
		params.Add("response_type", "code")

		authURL := "https://discord.com/oauth2/authorize" + params.Encode()
		fmt.Println("Opening browser to:", authURL)
		// TODO: "Press enter to open the following link in your browser"
		browser.OpenURL(authURL)

		// Block until we get a token or err
		select {
		case t := <-tokenChan:
			token = t
		case e := <-errChan:
			server.Shutdown(context.Background())
			return nil, e
		}

		server.Shutdown(context.Background())
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}
