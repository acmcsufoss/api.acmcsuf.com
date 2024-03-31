package resource

import (
	"context"
	"fmt"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/db/sqlite"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/services"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

type Resource struct {
	s services.Service
}

func NewResourceHandler(s services.Service) *Resource {
	return &Resource{s: s}
}

var _ services.Service = ResourcesService{}

type ResourcesService struct {
	q *sqlite.Queries
}

func New(q *sqlite.Queries) *ResourcesService {
	return &ResourcesService{q}
}

type resourceOutput sqlite.Resource

func (s ResourcesService) Resources() usecase.IOInteractor {
	res := s.q.GetResource(context.TODO(), "1")
	fmt.Println(res)
}

func (s ResourcesService) PostResources() usecase.IOInteractor {
	panic("implement me")
}

func (s ResourcesService) BatchPostResources() usecase.IOInteractor {
	panic("implement me")
}

func (s ResourcesService) GetResource() usecase.IOInteractor {
	// Create use case interactor with references to input/output types and interaction function.
	u := usecase.NewIOI(new(rdesourceInput), new(db.resourceOutput), func(ctx context.Context, input, output interface{}) error {
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
