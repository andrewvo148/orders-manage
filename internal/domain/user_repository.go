package domain

import (
	"context"
)

type UserRepository interface {
	All(ctx context.Context, offet int, limit int) ([]User, error)
	Save(ctx context.Context, order User) error
}
