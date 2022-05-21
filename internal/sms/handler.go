package sms

import (
	"net/http"

	"github.com/edsonmichaque/omni/internal"
)

type Handler struct {
	Cache internal.Cache
}

func (h Handler) SendMessages(rw http.ResponseWriter, r *http.Request) {}

func (h Handler) ListMessages(rw http.ResponseWriter, r *http.Request) {}

func (h Handler) GetMessage(rw http.ResponseWriter, r *http.Request) {}
