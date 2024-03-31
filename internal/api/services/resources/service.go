package resources

import (
	"context"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/db/sqlite"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/models"
)

type ResourceService struct {
	q *sqlite.Queries
}

func New(q *sqlite.Queries) *ResourceService {
	return &ResourceService{q}
}

func (s ResourceService) GetResource(uuid string) (*models.Resource, error) {
	res, err := s.q.GetResource(context.Background(), uuid)
	if err != nil {
		return nil, err
	}
	val := (models.Resource)(res)
	return &val, nil
}
