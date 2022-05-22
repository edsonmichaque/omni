package operator

import (
	"net/http"

	"github.com/edsonmichaque/omni/internal"
)

type Handler struct {
	internal.Handler
	Cache  internal.Cache
	Logger internal.Log
}

func (h Handler) SendMessages(rw http.ResponseWriter, r *http.Request) {
	var req NewSMS

	if err := h.ReadJSON(r.Body, &req); err != nil {
		return
	}
}

func (h Handler) ListMessages(rw http.ResponseWriter, r *http.Request) {}

func (h Handler) GetMessage(rw http.ResponseWriter, r *http.Request) {}
