package services

import (
	"encoding/json"
	"io"
	"os"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	disgoauth "github.com/realTristan/disgoauth"
)

var RequestClient *http.Client = &http.Client{}

type DiscordOauthServicer interface {
	Redirect(w gin.ResponseWriter, r *http.Request)
	HandleRedirect(code string)
}

type DiscordOauthService struct {
	dc *disgoauth.Client
}

func NewDiscordOauthService() *DiscordOauthService {
	return &DiscordOauthService{dc: disgoauth.Init(&disgoauth.Client{
								ClientID: os.Getenv("DISCORD_CLIENT_ID"),
								ClientSecret: os.Getenv("DISCORD_CLIENT_SECRET"),
								RedirectURI: os.Getenv("DISCORD_REDIRECT_URI"),
								Scopes: []string{disgoauth.ScopeIdentify, "guilds.members.read"},
	})}
}

func (s *DiscordOauthService) Redirect(w gin.ResponseWriter, r *http.Request) {
	s.dc.RedirectHandler(w, r, "")
}

func (s *DiscordOauthService) HandleRedirect(code string) {
	var (
		accessToken, _ = s.dc.GetOnlyAccessToken(code)
		guildData, _ = GetUserGuildData(accessToken)
	)
	jguildData, _ := json.MarshalIndent(guildData, "", "\t")
	fmt.Println(string(jguildData))
}

// Reused disgoauth.userData function and remapped to work
// in getting User's Guild Data, includes user info too.
func GetUserGuildData(token string) (map[string]any, error) {
	// Establish a new request object
	req, err := http.NewRequest("GET", "https://discord.com/api/users/@me/guilds/710225099923521558/member", nil)

	// Handle the error
	if err != nil {
		return map[string]any{}, err
	}
	// Set the request object's headers
	req.Header = http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{token},
	}
	// Send the http request
	resp, err := RequestClient.Do(req)

	// Handle the error
	// If the response status isn't a success
	if resp.StatusCode != 200 || err != nil {
		// Read the http body
		body, _err := io.ReadAll(resp.Body)

		// Handle the read body error
		if _err != nil {
			return map[string]any{}, _err
		}
		// Handle http response error
		return map[string]any{},
			fmt.Errorf("status: %d, code: %v, body: %s",
				resp.StatusCode, err, string(body))
	}

	// Readable golang map used for storing
	// the response body
	var data map[string]any

	// Handle the error
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return map[string]any{}, err
	}
	return data, nil
}
