package services

import (
	"context"

	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/db/sqlite"
)

var _ Service = ResourcesService{}

type ResourcesService struct {
	q *sqlite.Queries
}

func NewResourcesService(q *sqlite.Queries) *ResourcesService {
	return &ResourcesService{q}
}

type resourceInput struct {
	Title          string `json:"title"`
	ContentMd      string `json:"content_md"`
	ImageUrl       string `json:"image_url"`
	ResourceType   string `json:"resource_type"`
	ResourceListID string `json:"resource_list_id"`
}

type resourceOutput sqlite.Resource

func (s ResourcesService) Resources() usecase.IOInteractor {
	panic("implement me")
}

func (s ResourcesService) PostResources() usecase.IOInteractor {
	panic("implement me")
}

func (s ResourcesService) BatchPostResources() usecase.IOInteractor {
	panic("implement me")
}

func (s ResourcesService) Resource() usecase.IOInteractor {
	// Create use case interactor with references to input/output types and interaction function.
	u := usecase.NewIOI(new(resourceInput), new(resourceOutput), func(ctx context.Context, input, output interface{}) error {
		// var (
		// 	in  = input.(*resourceInput)
		// 	out = output.(*resourceOutput)
		// )

		// TODO: Get resource by ID from database.

		return nil
	})

	// Describe use case interactor.
	u.SetTitle("GetResource")
	u.SetDescription("Gets a single base resource.")
	u.SetExpectedErrors(status.InvalidArgument)
	return u
}

func (s ResourcesService) PostResource() usecase.IOInteractor {
	panic("implement me")
}

func (s ResourcesService) BatchPostResource() usecase.IOInteractor {
	panic("implement me")
}

func (s ResourcesService) DeleteResource() usecase.IOInteractor {
	panic("implement me")
}
