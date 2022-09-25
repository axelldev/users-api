package handlers

import (
	"net/http"
	"strconv"

	"github.com/axelldev/users-api/models"
	"github.com/gorilla/mux"
)

// GetUsersResponse to send a Message and an User slice.
type GetUsersResponse struct {
	Message string        `json:"message"`
	Users   []models.User `json:"users"`
}

// GetUserResponse to send a Message and a User instance.
type GetUserResponse struct {
	Message string      `json:"message"`
	User    models.User `json:"user"`
}

// GetUsers sends a json response with a slice of User.
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := RespondJSON(w, GetUsersResponse{
		Message: "success",
		Users:   []models.User{},
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GetUser responds with an User instance.
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	// Not found id on params
	if _, ok := params["id"]; !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Checks for a valid id
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := RespondJSON(w, GetUserResponse{
		Message: "success",
		User: models.User{
			ID:   id,
			Name: "Elon Mus",
			Age:  33,
		},
	}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
