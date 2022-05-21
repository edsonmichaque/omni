package ussd

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h Handler) Router(parent *mux.Router) *mux.Router {
	r := parent.NewRoute().Subrouter()

	r.HandleFunc("/ussd", h.SendMessages).Methods(http.MethodPost)
	r.HandleFunc("/ussd", h.ListMessages).Methods(http.MethodGet)
	r.HandleFunc("/ussd/:id", h.GetMessage).Methods(http.MethodGet)

	return r
}
