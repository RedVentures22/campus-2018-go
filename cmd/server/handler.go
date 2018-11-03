package main

import (
	"net/http"

	users "github.com/RedVentures22/campus-2018-go"
)

type handler struct {
	store *users.Store
}

func (h *handler) routerUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.listUsers(w, r)
	case http.MethodPost:
		h.createUser(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}

}
