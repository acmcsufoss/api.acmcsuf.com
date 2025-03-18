package services

import (
	"github.com/swaggest/usecase"
	"context"
	"database/sql"
	_ "embed"
	"log"
	"reflect"

	_ "modernc.org/sqlite"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/db/models"
)

type EventsService struct {
	q *models.Queries
	ctx context.Background
}

var ddl string

func NewEventsService(q *models.Queries) (*EventsService, error) {
	ctx := context.Background()

	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		return nil, err
	}

	// Create tables
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		return nil, err
	}

	queries := models.New(db)
	return &EventsService{queries}, nil
}


func GetEvent(q *models.Queries) (models.Event, error) {
	// I think this is the wrong way to implement since this only returns error
	// and passes around a context
	uuid := 1 // NOTE: This is a placeholder
	event, err := q.GetEvent(q.ctx, uuid)
	if err != nil {
		return nil, err
	}
	return event, nil
}

// TODO: re-implement this interface

// func (s EventsService) Resources() usecase.IOInteractor {
// 	panic("implement me")
// }
//
// func (s EventsService) PostResources() usecase.IOInteractor {
// 	panic("implement me")
// }
//
// func (s EventsService) BatchPostResources() usecase.IOInteractor {
// 	panic("implement me")
// }
//
// func (s EventsService) Resource() usecase.IOInteractor {
// 	panic("implement me")
// }
//
// func (s EventsService) PostResource() usecase.IOInteractor {
// 	panic("implement me")
// }
//
// func (s EventsService) BatchPostResource() usecase.IOInteractor {
// 	panic("implement me")
// }
//
// func (s EventsService) DeleteResource() usecase.IOInteractor {
// 	panic("implement me")
// }
