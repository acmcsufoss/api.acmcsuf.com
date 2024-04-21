package oauth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"io"
	"net/http"

	"github.com/ravener/discord-oauth2"
	"golang.org/x/oauth2"
)

// Utilized this resource to get started with the OAuth2 implementation:
// https://github.com/ravener/discord-oauth2/blob/master/example/main.go

var stateSize = 32
var state, stateErr = GenerateState(stateSize)

func CreateDiscordOAuth() {
	// How would we set up the secrets? .env file?
	conf := &oauth2.Config{
		RedirectURL: "http://localhost:3000/auth/callback",
		// This next 2 lines must be edited before running this.
		ClientID:     "id",
		ClientSecret: "secret",
		Scopes:       []string{discord.ScopeIdentify},
		Endpoint:     discord.Endpoint,
	}

	CreateOAuthEndpoints(*conf)
}

func CreateOAuthEndpoints(c oauth2.Config) {
	if stateErr != nil {
		panic("state is null")
	}

	// Login page that redirects to Discord's auth page
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, c.AuthCodeURL(state), http.StatusTemporaryRedirect)
	})

	// Callback endpoint that Discord redirects to
	http.HandleFunc("/auth/callback", func(w http.ResponseWriter, r *http.Request) {
		// Check if the state is valid
		if r.FormValue("state") != state {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("State does not match."))
			return
		}

		// Exchange the code for a token
		token, err := c.Exchange(context.Background(), r.FormValue("code"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		// Get a Discord user's info using the token
		res, err := c.Client(context.Background(), token).Get(GetDiscordUserEndpointURL())
		if err != nil || res.StatusCode != http.StatusOK {
			w.WriteHeader(http.StatusInternalServerError)
			errMsg := err.Error()
			if err == nil {
				errMsg = res.Status
			}
			w.Write([]byte(errMsg))
			return
		}

		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Write(body)
	})

}

// Generate a random string for OAuth2 state
// Source: https://stackoverflow.com/a/35558869
func GenerateState(n int) (string, error) {
	data := make([]byte, n)
	if _, err := io.ReadFull(rand.Reader, data); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

// Gets the URL for the Discord API user endpoint
func GetDiscordUserEndpointURL() string {
	return "https://discord.com/api/users/@me"
}
