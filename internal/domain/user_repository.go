package domain

import (
	"context"

	"github.com/andrewvo148/orders-manage/internal/domain"
)

type UserRepository interface {
	All(ctx context.Context, offet int, limit int) ([]domain.User, error)
	Save(ctx context.Context, order domain.User) error
}
