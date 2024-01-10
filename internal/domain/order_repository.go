package domain

import (
	"context"
)

type OrderRepository interface {
	All(ctx context.Context, offet int, limit int) ([]Order, error)
	Save(ctx context.Context, order Order) error
}
