package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func (api *Config) routes() http.Handler {
	router := chi.NewRouter()

	// specify who is allowed to connect
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Health check handlers
	router.Get("/health", api.handleHealthCheck)

	// Users handlers
	router.Post("/users", api.handleCreateUser)
	router.Get("/users", api.handleGetAllUsers)
	router.Get("/users/id/{id}", api.handleGetUserById)
	router.Get("/users/email/{email}", api.handleGetUserByEmail)
	router.Patch("/users/id/{id}", api.handleUpdateUserById)
	router.Delete("/users/id/{id}", api.handleDeleteUserById)

	return router
}
