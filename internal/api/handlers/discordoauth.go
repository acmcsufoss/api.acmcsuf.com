package handlers

import (
	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/services"
	"github.com/gin-gonic/gin"
)

type DiscordOauthHandler struct {
	discordOauthService services.DiscordOauthServicer
}

func NewDiscordOauthHandler(discordOauthService services.DiscordOauthServicer) *DiscordOauthHandler {
	return &DiscordOauthHandler{discordOauthService: discordOauthService}
}

func (h *DiscordOauthHandler) GoRedirect(c *gin.Context) {
	w := c.Writer
	r := c.Request
	h.discordOauthService.Redirect(w, r)

}

func (h *DiscordOauthHandler) HandleRedirect(c *gin.Context) {
	code := c.Request.URL.Query()["code"][0]
	h.discordOauthService.HandleRedirect(code)

}
