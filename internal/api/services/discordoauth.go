package services

import (
	"os"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	disgoauth "github.com/realTristan/disgoauth"
)

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
		userData, _ = disgoauth.GetUserData(accessToken)
	)
	fmt.Println(userData)
}
