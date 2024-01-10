package rest

import (
	"encoding/json"
	"github.com/andrewvo148/orders-manage/internal/application"
	"github.com/andrewvo148/orders-manage/internal/application/commands"
	"net/http"
)

type UserHandler struct {
	application.Application
}

func (h UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userRequest commands.CreateUserRequest

	// Decode JSON request body
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request body"))
		return
	}

	// Respond with the created user
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}
