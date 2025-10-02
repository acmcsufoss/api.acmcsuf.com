package services

import (
	"context"
)

type Service[T any, ID any, CreateParams any, UpdateParams any] interface {
	Get(ctx context.Context, id ID) (T, error)
	List(ctx context.Context, filters ...any) ([]T, error)
	Create(ctx context.Context, params CreateParams) error
	Update(ctx context.Context, id ID, params UpdateParams) error
	Delete(ctx context.Context, id ID) error
}
