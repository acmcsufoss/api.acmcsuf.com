package resource

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/db/sqlite"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/handlers"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/services"
)

type Resource struct {
	s services.Service
}

func NewResourceHandler(s services.Service) *ResourceHandler {
	return &ResourceHandler{s: s}
}

var _ handlers.Handler = ResourceHandler{}

type ResourceHandler struct {
	s services.Service
}

func New(s services.Service) *ResourceHandler {
	return &ResourceHandler{s}
}

type resourceOutput sqlite.Resource

// func (s ResourcesService) Resources() usecase.IOInteractor {
// 	res := s.q.GetResource(context.TODO(), "1")
// 	fmt.Println(res)
// }

// func (s ResourcesService) PostResources() usecase.IOInteractor {
// 	panic("implement me")
// }

// func (s ResourcesService) BatchPostResources() usecase.IOInteractor {
// 	panic("implement me")
// }

// func (s ResourceHandler) GetResource(w http.ResponseWriter, r *http.Request) usecase.IOInteractor {

// 	// Create use case interactor with references to input/output types and interaction function.
// 	u := usecase.NewIOI(new(rdesourceInput), new(db.resourceOutput), func(ctx context.Context, input, output interface{}) error {
// 		// var (
// 		// 	in  = input.(*resourceInput)
// 		// 	out = output.(*resourceOutput)
// 		// )

// 		// TODO: Get resource by ID from database.

// 		return nil
// 	})

// 	// Describe use case interactor.
// 	u.SetTitle("GetResource")
// 	u.SetDescription("Gets a single base resource.")
// 	u.SetExpectedErrors(status.InvalidArgument)
// 	return u
// }

func (re ResourceHandler) GetResource(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	res, err := re.s.GetResource(context.Background(), id)
	if err != nil {
		log.Println("There was an error processing your request. %v", err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// func (s ResourcesService) PostResource() usecase.IOInteractor {
// 	panic("implement me")
// }

// func (s ResourcesService) BatchPostResource() usecase.IOInteractor {
// 	panic("implement me")
// }

// func (s ResourcesService) DeleteResource() usecase.IOInteractor {
// 	panic("implement me")
// }
