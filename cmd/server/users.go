package main

import (
	"encoding/json"
	"net/http"

	users "github.com/RedVentures22/campus-2018-go"
)

type createUserRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type createUserResponse struct {
	ID int64 `json:"id"`
	createUserRequest
}

func (h *handler) createUser(w http.ResponseWriter, r *http.Request) {
	var req createUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u := users.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	id := h.store.AddUser(u)

	resp := createUserResponse{
		ID:                id,
		createUserRequest: req,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

type listUsersResponse struct {
	Users []users.User `json:"users"`
}

func (h *handler) listUsers(w http.ResponseWriter, r *http.Request) {
	users := h.store.ListUsers()

	resp := listUsersResponse{
		Users: users,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
