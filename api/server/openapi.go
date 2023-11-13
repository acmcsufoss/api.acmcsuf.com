package server

// Copied from:
// https://pkg.go.dev/github.com/swaggest/rest@v0.2.59/web#example-DefaultService
//
// TODO: Replace with our own implementation.

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	"github.com/swaggest/openapi-go/openapi3"
	"github.com/swaggest/rest/nethttp"
	"github.com/swaggest/rest/web"
	"github.com/swaggest/usecase"
)

// album represents data about a record album.
type album struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
	Locale string  `query:"locale"`
}

func postAlbums() usecase.Interactor {
	u := usecase.NewIOI(new(album), new(album), func(ctx context.Context, input, output interface{}) error {
		log.Println("Creating album")

		return nil
	})
	u.SetTags("Album")

	return u
}

func NewOpenAPI() http.Handler {
	// Service initializes router with required middlewares.
	service := web.NewService(openapi3.NewReflector())

	// It allows OpenAPI configuration.
	service.OpenAPISchema().SetTitle("api.acmcsuf.com")
	service.OpenAPISchema().SetDescription("This is the API server for api.acmcsuf.com.")
	service.OpenAPISchema().SetVersion("0.0.1-unreleased")

	// Additional middlewares can be added.
	service.Use(
		middleware.StripSlashes,

		cors.AllowAll().Handler, // "github.com/rs/cors", 3rd-party CORS middleware can also be configured here.
	)

	service.Wrap()

	// Use cases can be mounted using short syntax .<Method>(...).
	service.Post("/albums", postAlbums(), nethttp.SuccessStatus(http.StatusCreated))

	return service
}
