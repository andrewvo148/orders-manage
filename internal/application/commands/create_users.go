package commands

import (
	"context"

	"github.com/andrewvo148/orders-manage/internal/domain"
)

type CreateUserHandler struct {
	userRepo domain.UserRepository
}

type CreateUserRequest struct {
	ID       string
	Fullname string
	Phone    string
	Email    string
	Password string
}

func NewCreateUserHandler(userRepo domain.UserRepository) *CreateUserHandler {
	return &CreateUserHandler{userRepo: userRepo}
}

func (h CreateUserHandler) CreateUsers(ctx context.Context, request CreateUserRequest) error {
	return h.userRepo.Save(ctx, domain.User{
		ID:       request.ID,
		FullName: request.Fullname,
		Email:    request.Email,
		Phone:    request.Phone,
	})
}
