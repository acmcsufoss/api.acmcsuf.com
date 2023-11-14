package server

// Copied from:
// https://pkg.go.dev/github.com/swaggest/rest@v0.2.59/web#example-DefaultService
//
// TODO: Replace with our own implementation.

import (
	"context"
	"net/http"

	"github.com/acmcsufoss/api.acmcsuf.com/api"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	"github.com/swaggest/openapi-go/openapi3"
	"github.com/swaggest/rest/nethttp"
	"github.com/swaggest/rest/web"
	"github.com/swaggest/usecase"
)

func postEvents(s api.Store) usecase.Interactor {
	return usecase.NewInteractor(func(ctx context.Context, input api.CreateEventRequest, _ *interface{}) error {
		_, err := s.CreateEvent(input)
		if err != nil {
			return err
		}

		return nil
	})
}

func NewOpenAPI(s api.Store) http.Handler {
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
	service.Get("/events", postEvents(s), nethttp.SuccessStatus(http.StatusCreated))
	service.Post("/events", postEvents(s), nethttp.SuccessStatus(http.StatusCreated))
	// service.Get("/events/{id}", getEvent(), nethttp.SuccessStatus(http.StatusOK))

	return service
}
