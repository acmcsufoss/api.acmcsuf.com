package services

import (
	"github.com/swaggest/usecase"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/db/models"
)

type EventsService struct {
	q *models.Queries
}

func NewEventsService(q *models.Queries) *EventsService {
	return &EventsService{q}
}

func GetEvent(q *models.Queries) models.Event {
	// I think this is the wrong way to implement since this only returns error
	// and passes around a context
	// error := q.GetEvent(ctx context.Context, uuid string)
	return models.Event{}
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
