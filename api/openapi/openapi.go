package openapi

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
	service.OpenAPISchema().SetDescription("This is api.acmcsuf.com, a data layer for acmcsuf.com.")
	service.OpenAPISchema().SetVersion("0.0.1-unreleased")

	// Additional middlewares can be added.
	service.Use(
		middleware.StripSlashes,
		cors.AllowAll().Handler,
	)
	service.Wrap()

	// crud(service, "/resource-lists", postEvents(s), nil, nil, nil, nil)
	crud(service, "/events", postEvents(s), nil, nil, nil, nil)
	// crud(service, "/announcements", postEvents(s), nil, nil, nil, nil)
	// crud(service, "/blog-posts", createBlogPost(s), readBlogPost(s), updateBlogPost(s), deleteBlogPost(s), listBlogPosts(s))

	return service
}

func crud(service *web.Service, patternPrefix string, creater, reader, updater, deleter, lister usecase.Interactor) {
	if creater != nil {
		service.Post(patternPrefix, creater, nethttp.SuccessStatus(http.StatusCreated))
	}

	if reader != nil {
		service.Get(patternPrefix+"/{id}", reader, nethttp.SuccessStatus(http.StatusOK))
	}

	if updater != nil {
		service.Post(patternPrefix+"/{id}", updater, nethttp.SuccessStatus(http.StatusOK))
	}

	if deleter != nil {
		service.Delete(patternPrefix+"/{id}", deleter, nethttp.SuccessStatus(http.StatusOK))
	}

	if lister != nil {
		service.Get(patternPrefix, lister, nethttp.SuccessStatus(http.StatusOK))
	}
}
