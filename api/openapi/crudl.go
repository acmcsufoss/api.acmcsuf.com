package openapi

import (
	"net/http"

	"github.com/swaggest/rest/nethttp"
	"github.com/swaggest/rest/web"
	"github.com/swaggest/usecase"
)

// crudl is a helper function for registering create, read, update, delete,
// and list usecases.
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
