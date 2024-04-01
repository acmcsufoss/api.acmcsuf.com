package handlers

import "net/http"

// Service is the interface of API endpoints for a resource service.
type Handler interface {
	// Resources gets a list of paginated resource resources.
	// Resources() usecase.IOInteractor

	// // PostResources creates a new resource resource.
	// PostResources() usecase.IOInteractor

	// // BatchPostResources creates multiple new resource resources.
	// BatchPostResources() usecase.IOInteractor

	// Resource gets a single resource resource.
	GetResource(w http.ResponseWriter, r *http.Request)
	// PostResource updates a single resource resource.
	// PostResource() usecase.IOInteractor

	// // BatchPostResource updates multiple resource resources.
	// BatchPostResource() usecase.IOInteractor

	// // DeleteResource deletes a single resource resource.
	// DeleteResource() usecase.IOInteractor
}
