package api

import (
	"github.com/swaggest/openapi-go/openapi3"
	"github.com/swaggest/rest/response/gzip"
	"github.com/swaggest/rest/web"
	swgui "github.com/swaggest/swgui/v5emb"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/services"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/services/events"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/services/resources"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/db/sqlite"
)

func New(q *sqlite.Queries) *web.Service {
	s := web.NewService(openapi3.NewReflector())

	// Init API documentation schema.
	s.OpenAPISchema().SetTitle("Basic Example")
	s.OpenAPISchema().SetDescription("This app showcases a trivial REST API.")
	s.OpenAPISchema().SetVersion("v0.0.1")

	// Setup middlewares.
	s.Wrap(
		gzip.Middleware, // Response compression with support for direct gzip pass through.
	)

	// Add use case handler to router.
	// s.Get("/hello/{name}", helloWorld())
	useAll(s, q)

	// Swagger UI endpoint at /docs.
	s.Docs("/docs", swgui.New)

	return s
}

func useAll(s *web.Service, q *sqlite.Queries) {
	use("/resources", resources.New(q), s)
	use("/events", events.New(q), s)
}

func use(path string, s services.Service, ss *web.Service) {
	ss.Get(path, s.Resources())
	ss.Post(path, s.PostResources())
	ss.Post(path, s.BatchPostResources())
	ss.Get(path+"/{id}", s.Resource())
	ss.Post(path+"/{id}", s.PostResource())
	ss.Post(path+"/{id}", s.BatchPostResource())
	ss.Delete(path+"/{id}", s.DeleteResource())
}
