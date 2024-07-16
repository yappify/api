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

type CreateUserPayload struct {
	AuthType string `json:"auth_type"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

type UpdateUserPayload struct {
	AuthType string `json:"auth_type,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	IsBanned bool   `json:"is_banned,omitempty"`
}

type Message struct {
	AuthorID       string `json:"authorID"`
	Username       string `json:"username"`
	RoomID         string `json:"roomID"`
	MessageContent string `json:"messageContent"`
	Timestamp      string `json:"timestamp"`
}
