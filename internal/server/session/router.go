package session

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h Handler) Router(r *mux.Router) *mux.Router {
	r.HandleFunc("/ussd", h.SendMessages).Methods(http.MethodPost)
	r.HandleFunc("/ussd", h.ListMessages).Methods(http.MethodGet)
	r.HandleFunc("/ussd/:id", h.GetMessage).Methods(http.MethodGet)

	return r
}
