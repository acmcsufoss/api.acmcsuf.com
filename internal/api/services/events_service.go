package services

import (
	"github.com/swaggest/usecase"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/db/sqlite"
)

var _ Service = EventsService{}

type EventsService struct {
	q *sqlite.Queries
}

func NewEventsService(q *sqlite.Queries) *EventsService {
	return &EventsService{q}
}

func (s EventsService) Resources() usecase.IOInteractor {
	panic("implement me")
}

func (s EventsService) PostResources() usecase.IOInteractor {
	panic("implement me")
}

func (s EventsService) BatchPostResources() usecase.IOInteractor {
	panic("implement me")
}

func (s EventsService) Resource() usecase.IOInteractor {
	panic("implement me")
}

func (s EventsService) PostResource() usecase.IOInteractor {
	panic("implement me")
}

func (s EventsService) BatchPostResource() usecase.IOInteractor {
	panic("implement me")
}

func (s EventsService) DeleteResource() usecase.IOInteractor {
	panic("implement me")
}
