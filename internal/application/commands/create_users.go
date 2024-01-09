package commands

import (
	"context"

	"github.com/andrewvo148/orders-manage/internal/domain"
)

type CreateUserHandler struct {
	userRepo domain.UserRepository
}

type CreateUserRequest struct {
	Fullname string
	Phone    string
	Email    string
	Password string
}

func (h CreateUserHandler) CreateUsers(ctx context.Context, request CreateUserRequest) {
	h.userRepo.Save(ctc, )
}
