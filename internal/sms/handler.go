package sms

import "net/http"

type Handler struct{}

func (h Handler) SendMessages(rw http.ResponseWriter, r *http.Request) {}

func (h Handler) ListMessages(rw http.ResponseWriter, r *http.Request) {}

func (h Handler) GetMessage(rw http.ResponseWriter, r *http.Request) {}
