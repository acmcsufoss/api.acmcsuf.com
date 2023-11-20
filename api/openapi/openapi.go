package openapi

import (
	"context"
	"net/http"

	"github.com/acmcsufoss/api.acmcsuf.com/api"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	"github.com/swaggest/openapi-go/openapi3"
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
	useCRUDL(service, withPrefix("/events"), withCreater(postEvents(s)))
	// crud(service, "/announcements", postEvents(s), nil, nil, nil, nil)
	// crud(service, "/blog-posts", createBlogPost(s), readBlogPost(s), updateBlogPost(s), deleteBlogPost(s), listBlogPosts(s))

	return service
}
