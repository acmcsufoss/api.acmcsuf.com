package requests

import (
	"net/http"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/config"
)

func AddOrigin(r *http.Request) {
	org := config.Cfg.Origin

	r.Header.Set("Origin", org)
}
