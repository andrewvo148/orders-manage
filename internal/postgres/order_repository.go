package postgres

import (
	"context"

	"github.com/andrewvo148/orders-manage/internal/domain"
)

type OrderRepository struct {
}

func All(ctx context.Context, offset int, limit int) ([]domain.Order, error) {

}
