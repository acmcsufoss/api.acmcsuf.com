package services

import "context"

// Service is the interface of API endpoints for a resource service.
type Service interface {
	// // Resources gets a list of paginated resource resources.
	// Resources() usecase.IOInteractor

	// // PostResources creates a new resource resource.
	// PostResources() usecase.IOInteractor

	// // BatchPostResources creates multiple new resource resources.
	// BatchPostResources() usecase.IOInteractor

	// Resource gets a single resource resource.
	GetResource(ctx context.Context,uuid string) (*models.Resource, error) 

// 	// PostResource updates a single resource resource.
// 	PostResource() usecase.IOInteractor

// 	// BatchPostResource updates multiple resource resources.
// 	BatchPostResource() usecase.IOInteractor

// 	// DeleteResource deletes a single resource resource.
// 	DeleteResource() usecase.IOInteractor
// }
