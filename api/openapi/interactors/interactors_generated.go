// Code is generated. DO NOT EDIT.
package interactors

import (
	"context"
	"net/http"

	"github.com/swaggest/rest/nethttp"
	"github.com/swaggest/rest/web"
	"github.com/swaggest/usecase"

	"github.com/acmcsufoss/api.acmcsuf.com/api"
)

// crudl is a helper function for registering create, read, update, delete,
// and list usecases for a given resource.
type crudl struct {
	patternPrefix string
	creater       usecase.Interactor
	reader        usecase.Interactor
	updater       usecase.Interactor
	deleter       usecase.Interactor
	lister        usecase.Interactor
}

type crudlOptionFn func(*crudl)

func withPrefix(prefix string) crudlOptionFn {
	return func(o *crudl) {
		o.patternPrefix = prefix
	}
}

func withCreate(creater usecase.Interactor) crudlOptionFn {
	return func(o *crudl) {
		o.creater = creater
	}
}

func withRead(reader usecase.Interactor) crudlOptionFn {
	return func(o *crudl) {
		o.reader = reader
	}
}

func withUpdate(updater usecase.Interactor) crudlOptionFn {
	return func(o *crudl) {
		o.updater = updater
	}
}

func withDelete(deleter usecase.Interactor) crudlOptionFn {
	return func(o *crudl) {
		o.deleter = deleter
	}
}

func withList(lister usecase.Interactor) crudlOptionFn {
	return func(o *crudl) {
		o.lister = lister
	}
}

func useCRUDL(service *web.Service, optionFns ...crudlOptionFn) {
	o := crudl{}
	for _, fn := range optionFns {
		fn(&o)
	}

	if o.creater != nil {
		service.Post(o.patternPrefix, o.creater, nethttp.SuccessStatus(http.StatusCreated))
	}

	if o.reader != nil {
		service.Get(o.patternPrefix+"/{id}", o.reader, nethttp.SuccessStatus(http.StatusOK))
	}

	if o.updater != nil {
		service.Post(o.patternPrefix+"/{id}", o.updater, nethttp.SuccessStatus(http.StatusOK))
	}

	if o.deleter != nil {
		service.Delete(o.patternPrefix+"/{id}", o.deleter, nethttp.SuccessStatus(http.StatusOK))
	}

	if o.lister != nil {
		service.Get(o.patternPrefix, o.lister, nethttp.SuccessStatus(http.StatusOK))
	}
}

// UseEvent uses a generated Event interactor.
func UseEvents(service *web.Service, store api.Store) {
	useCRUDL(
		service,
		withPrefix("/events"),
		withCreate(usecase.NewInteractor(func(ctx context.Context, request *api.CreateEventRequest, response *api.CreateEventResponse) (err error) {
			*response, err = store.CreateEvent(request)
			return err
		})),
		withRead(usecase.NewInteractor(func(ctx context.Context, request *api.ReadEventRequest, response *api.ReadEventResponse) (err error) {
			*response, err = store.ReadEvent(request)
			return err
		})),
		withUpdate(usecase.NewInteractor(func(ctx context.Context, request *api.UpdateEventRequest, response *api.UpdateEventResponse) (err error) {
			*response, err = store.UpdateEvent(request)
			return err
		})),
		withDelete(usecase.NewInteractor(func(ctx context.Context, request *api.DeleteEventRequest, response *api.DeleteEventResponse) (err error) {
			*response, err = store.DeleteEvent(request)
			return err
		})),
		withList(usecase.NewInteractor(func(ctx context.Context, request *api.ListEventRequest, response *api.ListEventResponse) (err error) {
			*response, err = store.ListEvent(request)
			return err
		})),
	)
}

// UseAll uses all interactors.
func UseAll(service *web.Service, store api.Store) {
	UseEvents(service, store)
}
