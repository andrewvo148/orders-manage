package postgres

import (
	"context"
	"github.com/andrewvo148/orders-manage/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	poolConn *pgxpool.Pool
}

func NewUserRepository(poolConn *pgxpool.Pool) *UserRepository {
	return &UserRepository{poolConn: poolConn}
}

func (userRepo UserRepository) Save(ctx context.Context, order domain.User) error {
	_, err := userRepo.poolConn.Exec(ctx, "INSERT INTO users(id, fullname, email, phone) VALUES ($1, $2, $3, $4)",
		order.ID, order.FullName, order.Email, order.Phone)
	return err
}
