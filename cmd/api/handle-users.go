package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/yappify/api/internal/database"
)

// [POST] handleCreateUser: create a new user
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
func (api *Config) handleGetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := api.DB.GetAllUsers(r.Context())

	if err != nil {
		api.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	api.writeJSON(w, http.StatusOK, users)
}

// [GET] handleGetUserById: fetch user by id from database
func (api *Config) handleGetUserById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		api.errorJSON(w, fmt.Errorf("invalid UUID format"), http.StatusBadRequest)
		return
	}

	user, err := api.DB.GetUserById(r.Context(), id)

	if err != nil {
		api.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	api.writeJSON(w, http.StatusOK, user)
}

// [GET] handleGetUserByEmail: fetch user by email from database
func (api *Config) handleGetUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")

	user, err := api.DB.GetUserByEmail(r.Context(), email)

	if err != nil {
		api.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	api.writeJSON(w, http.StatusOK, user)
}

// [PATCH] handleUpdateUserById: fetch user by id from database and update
func (api *Config) handleUpdateUserById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		api.errorJSON(w, fmt.Errorf("invalid UUID format"), http.StatusBadRequest)
		return
	}

	var payload UpdateUserPayload

	if err := api.readJSON(w, r, &payload); err != nil {
		api.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	user, err := api.DB.UpdateUserById(r.Context(), database.UpdateUserByIdParams{
		ID:       id,
		AuthType: payload.AuthType,
		Name:     payload.Name,
		Email:    payload.Email,
		Password: sql.NullString{String: payload.Password, Valid: true},
		IsBanned: payload.IsBanned,
	})

	if err != nil {
		api.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	api.writeJSON(w, http.StatusCreated, user)
}

// [DELETE] handleDeleteUserById: delete user by id from database
func (api *Config) handleDeleteUserById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		api.errorJSON(w, fmt.Errorf("invalid UUID format"), http.StatusBadRequest)
		return
	}

	err = api.DB.DeleteUserById(r.Context(), id)

	if err != nil {
		api.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	var payload JsonResponse
	payload.Error = false
	payload.Message = fmt.Sprintf("Successfully deleted user of id %s", id)

	api.writeJSON(w, http.StatusOK, payload)
}
