package main

import "github.com/yappify/api/internal/database"

type Config struct {
	DB *database.Queries
}

type JsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// id, auth_type, name, email, password, is_banned
type CreateUserPayload struct {
	AuthType string `json:"auth_type"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}
