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
	"path/filepath"
	"time"

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

type StoredToken struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	Expiry       time.Time `json:"expiry"`
}

// NOTE: As far as I can tell this port must be hardcoded as it needs to match the port
// specified in the discord developer portal. It's chosen arbitrarily and is hopefully random
// enough to not run into port conflicts.
const redirectAddr = ":61234"

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
		if storedToken, err := loadToken(); err == nil {
			token = storedToken
		} else {
			fmt.Println("No valid session found. Authenticating...")

			clientID := os.Getenv("DISCORD_CLIENT_ID")
			if clientID == "" {
				return nil, errors.New("DISCORD_CLIENT_ID is unset")
			}
			clientSecret := os.Getenv("DISCORD_CLIENT_SECRET")
			if clientSecret == "" {
				return nil, errors.New("DISCORD_CLIENT_SECRET is unset")
			}

			redirectURI := fmt.Sprintf("http://localhost%s", redirectAddr)
			const scope = "identify"

			tokenChan := make(chan string)
			errChan := make(chan error)
			mux := http.NewServeMux()
			server := &http.Server{Addr: redirectAddr, Handler: mux}
			mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				code := r.URL.Query().Get("code")
				if code == "" {
					fmt.Fprintln(w, "No code found")
					return
				}
				fmt.Fprintf(w, "Got code! You can close this window.")

				data := url.Values{}
				data.Set("client_id", clientID)
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

				// save token to disk HERE >:)
				if err := saveToken(tokenResp); err != nil {
					fmt.Printf("Warning: failed to save auth token: %v\n", err)
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

			baseURL := "https://discord.com/oauth2/authorize"
			u, _ := url.Parse(baseURL)
			u.RawQuery = params.Encode()
			fmt.Println("Opening browser to:", u.String())
			// TODO: "Press enter to open the following link in your browser"
			browser.OpenURL(u.String())

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
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

// ============== persistence helper functions ==============

func getTokenPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	// '~/.config/acmcsuf-cli/token.json' on Unix systems
	appDir := filepath.Join(configDir, "acmcsuf-cli")
	if err := os.MkdirAll(appDir, 0700); err != nil {
		return "", err
	}
	return filepath.Join(appDir, "token.json"), nil
}

func saveToken(resp TokenResponse) error {
	path, err := getTokenPath()
	if err != nil {
		return err
	}

	stored := StoredToken{
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
		Expiry:       time.Now().Add(time.Duration(resp.ExpiresIn-10) * time.Second),
	}

	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(stored)
}

func loadToken() (string, error) {
	path, err := getTokenPath()
	if err != nil {
		return "", err
	}

	file, err := os.Open(path)
	if err != nil {
		return "", err
	}

	var stored StoredToken
	if err := json.NewDecoder(file).Decode(&stored); err != nil {
		return "", err
	}

	if time.Now().After(stored.Expiry) {
		return "", errors.New("token expired")
	}

	return stored.AccessToken, nil
}
