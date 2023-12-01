// Code is generated. DO NOT EDIT.

package interactors

import (
	"context"
	"io"
	"net/http"

	"github.com/swaggest/rest/nethttp"
	"github.com/swaggest/rest/web"
	"github.com/swaggest/usecase"
)

// crud is a helper struct for registering create, read, update, delete,
// and batch usecases for a given resource.
type crud struct {
	patternPrefix string
	creater       usecase.Interactor
	batchCreater  usecase.Interactor
	reader        usecase.Interactor
	batchReader   usecase.Interactor
	updater       usecase.Interactor
	batchUpdater  usecase.Interactor
	deleter       usecase.Interactor
	batchDeleter  usecase.Interactor
}

// TODO: Replace instances of 'crudl' with 'crud'.

// TODO: Split up the functions into templates so they are included in the
// generated code only if they are used.

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

// ContainsContext can be embedded by any interface to have an overrideable
// context.
type ContainsContext interface {
	WithContext(context.Context) ContainsContext
}

// Store is the store interface.
type Store interface {
	io.Closer
	ContainsContext

}

// UseEvents uses a generated Events interactor.
func UseEvents(service *web.Service, store Store) {
	useCRUDL(
		service,
		withPrefix("/events"),
	)
}

// UseAll uses all interactors.
func UseAll(service *web.Service, store Store) {
	UseEvents(service, store)
}
