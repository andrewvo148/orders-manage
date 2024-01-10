package application

import "github.com/andrewvo148/orders-manage/internal/application/commands"

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateUser commands.CreateUserHandler
	//CreateUsers(ctx context.Context, request commands.CreateUserRequest) error
}

type Queries struct {
}
