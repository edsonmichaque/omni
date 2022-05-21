package session

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h Handler) Router(parent *mux.Router) *mux.Router {
	r := parent.NewRoute().Subrouter()

	r.HandleFunc("/messages", h.SendMessages).Methods(http.MethodPost)
	r.HandleFunc("/messages", h.ListMessages).Methods(http.MethodGet)
	r.HandleFunc("/messages/:id", h.GetMessage).Methods(http.MethodGet)

	return r
}
