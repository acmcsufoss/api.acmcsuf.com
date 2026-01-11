package requests

import (
	"net/http"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/config"
)

func AddOrigin(r *http.Request) {
	cfg := config.Load()
	org := cfg.AllowedOrigins[0]

	r.Header.Set("Origin", org)
}
