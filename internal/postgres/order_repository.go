package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepository struct {
	poolConn *pgxpool.Pool
}

func NewOrderRepository(poolConn *pgxpool.Pool) *OrderRepository {
	return &OrderRepository{poolConn: poolConn}
}
