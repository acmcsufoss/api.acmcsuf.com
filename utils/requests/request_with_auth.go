package requests

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

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
		clientID := os.Getenv("DISCORD_CLIENT_ID")
		if clientID == "" {
			return nil, errors.New("DISCORD_CLIENT_ID is unset")
		}
		// TODO: check that this port isn't being used first
		const redirectURI = "http://localhost:8888"
		const scope = "identify"
		params := url.Values{}
		params.Add("client_id", clientID)
		params.Add("redirect_uri", redirectURI)
		params.Add("scope", scope)
		params.Add("response_type", "code")

		authURL := "https://discord.com/oauth2/authorize" + params.Encode()
		fmt.Println("Opening browser to:", authURL)
		browser.OpenURL(authURL)
		// TODO: "Press enter to open the following link in your browser"
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			code := r.URL.Query().Get("code")
			if code != "" {
				fmt.Fprintf(w, "Got code! You can close this window.\n\nCode: %s", code)
				fmt.Printf("Success! Auth code: %s\n", code)

				fmt.Println("Exchanging code for access token...")
				data := url.Values{}
				data.set("client_id", clientID)
				clientSecret := os.Getenv("DISCORD_CLIENT_SECRET")
				if clientSecret == "" {
					fmt.Fprintf(os.Stderr, "DISCORD_CLIENT_SECRET is unset")
					return
				}
				data.set("client_secret", clientSecret)
				data.set("grant_type", "authorization_code")
				data.set("code", code)
				data.set("redirectURI", redirectURI)

				req, _ := http.NewRequest("POST", "https://discord.com/api/oauth2/token",
					strings.NewReader(data.Encode()))
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

				client := &http.Client{}
				resp, err := client.Do(req)
				if err != nil {
					fmt.Printf("Error exchanging token: %v\n", err)
					return
				}

				defer resp.Body.Close()
				body, _ := io.ReadAll(resp.Body)

				var tokenResp TokenResponse
				if err := json.Unmarshal(body, &tokenResp); err != nil {
					fmt.Printf("Error parsing JSON: %v\nResponse Body: %s\n", err, string(body))
					return
				}

				if tokenResp.AccessToken != "" {
					token = tokenResp.AccessToken
				} else {
					fmt.Fprintf(os.Stderr, "Error: Failed to get token. Discord said:\n%s\n", string(body))
				}

				go func() {
					return
				}()
			} else {
				fmt.Fprintln(os.Stderr, "Error: no code in URL")
			}
		})

		http.ListenAndServe(":8888", nil)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}
