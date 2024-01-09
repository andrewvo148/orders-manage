package rest

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Run() *chi.Mux {
	// Create a new Chi router
	r := chi.NewRouter()
	// Use built-in middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	
	return r
}
