package repository

import (
	"context"
)

type Repository[T any, ID any, Update any] interface {
	GetAll(ctx context.Context) ([]T, error)

	GetByID(ctx context.Context, id ID) (T, error)
	Delete(ctx context.Context, id ID) error

	Create(ctx context.Context, args T) error
	Update(ctx context.Context, args Update) error
}
