package openapi

import (
	"net/http"

	"github.com/acmcsufoss/api.acmcsuf.com/api"
	"github.com/acmcsufoss/api.acmcsuf.com/api/openapi/interactors"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	"github.com/swaggest/openapi-go/openapi3"
	"github.com/swaggest/rest/web"
)

func NewOpenAPI(s api.Store) http.Handler {
	// Service initializes router with required middlewares.
	service := web.NewService(openapi3.NewReflector())

	// It allows OpenAPI configuration.
	service.OpenAPISchema().SetTitle("api.acmcsuf.com")
	service.OpenAPISchema().SetDescription("This is api.acmcsuf.com, a data layer for acmcsuf.com.")
	service.OpenAPISchema().SetVersion("0.0.1-unreleased")

	// Additional middlewares can be added.
	service.Use(
		middleware.StripSlashes,
		cors.AllowAll().Handler,
	)
	service.Wrap()

	// Register API handler interactors with the service.
	interactors.UseAll(service, s)

	return service
}
