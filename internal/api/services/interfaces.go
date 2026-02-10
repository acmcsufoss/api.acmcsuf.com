package services

import (
	"context"
)

type Service[T any, ID any] interface {
	Get(ctx context.Context, id ID) (T, error)
	List(ctx context.Context, filters ...any) ([]T, error)
	Create(ctx context.Context, params T) error
	Update(ctx context.Context, id ID, params T) error
	Delete(ctx context.Context, id ID) error
}
