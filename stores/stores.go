package stores

//go:generate sqlc generate

import (
	"io"
	"log"

	"github.com/acmcsufoss/api.acmcsuf.com/api"
)

type StoreCloser interface {
	io.Closer
	api.Store
}

func Must[T api.Store](store T, err error) T {
	if err != nil {
		log.Fatalln("cannot make store:", err)
	}

	return store
}
