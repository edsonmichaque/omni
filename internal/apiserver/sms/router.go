package sms

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h Handler) Router(parent *mux.Router) *mux.Router {
	r := parent.NewRoute().Subrouter()

	r.HandleFunc("sms/", h.SendMessages).Methods(http.MethodPost)
	r.HandleFunc("sms/", h.ListMessages).Methods(http.MethodGet)
	r.HandleFunc("sms/:id", h.GetMessage).Methods(http.MethodGet)

	return r
}
