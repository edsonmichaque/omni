package operator

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h Handler) Router(r *mux.Router) {
	r.HandleFunc("/sms", h.SendMessages).Methods(http.MethodPost)
	r.HandleFunc("/sms", h.ListMessages).Methods(http.MethodGet)
	r.HandleFunc("/sms/:id", h.GetMessage).Methods(http.MethodGet)
}
