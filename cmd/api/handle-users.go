package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/yappify/api/internal/database"
)

// [POST] handleCreateUser: create a new user
//
//	type CreateUserPayload struct {
//		AuthType string `json:"auth_type"`
//		Name     string `json:"name"`
//		Email    string `json:"email"`
//		Password string `json:"password,omitempty"`
//	}
func (api *Config) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var payload CreateUserPayload

	if err := api.readJSON(w, r, &payload); err != nil {
		api.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	if err := validateCreateUserPayload(payload); err != nil {
		api.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	user, err := api.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		AuthType:  payload.AuthType,
		Name:      payload.Name,
		Email:     payload.Email,
		Password:  sql.NullString{String: payload.Password, Valid: true},
		IsBanned:  false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		api.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	api.writeJSON(w, http.StatusCreated, user)
}

// [GET] handleGetAllUsers: fetch all users from database
func (api *Config) handleGetAllUsers() {

}

// [GET] handleGetUserById: fetch user by id from database
func (api *Config) handleGetUserById() {

}

// [GET] handleGetUserByEmail: fetch user by email from database
func (api *Config) handleGetUserByEmail() {

}

// [PATCH] handleUpdateUserById: fetch user by id from database and update
func (api *Config) handleUpdateUserById() {

}

// [DELETE] handleDeleteUserById: delete user by id from database
func (api *Config) handleDeleteUserById() {

}
