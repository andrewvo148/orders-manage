package domain

import (
	"context"

	"github.com/andrewvo148/orders-manage/internal/domain"
)

type OrderRepository interface {
	All(ctx context.Context, offet int, limit int) ([]domain.Order, error)
	Save(ctx context.Context, order domain.Order) error
}
