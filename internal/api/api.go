package api

import (
	"github.com/swaggest/openapi-go/openapi3"
	"github.com/swaggest/rest/response/gzip"
	"github.com/swaggest/rest/web"
	swgui "github.com/swaggest/swgui/v5emb"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/db/models"
)

// func New(q *sqlite.Queries) *web.Service {
// 	s := web.NewService(openapi3.NewReflector())
//
// 	// Init API documentation schema.
// 	s.OpenAPISchema().SetTitle("Basic Example")
// 	s.OpenAPISchema().SetDescription("This app showcases a trivial REST API.")
// 	s.OpenAPISchema().SetVersion("v0.0.1")
//
// 	// Setup middlewares.
// 	s.Wrap(
// 		gzip.Middleware, // Response compression with support for direct gzip pass through.
// 	)
//
// 	// Swagger UI endpoint at /docs.
// 	s.Docs("/docs", swgui.New)
//
// 	return s
// }
