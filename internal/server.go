package internal

import (
	"net/http"

	"github.com/edsonmichaque/omni/internal/sms"
	"github.com/edsonmichaque/omni/internal/ussd"
	"github.com/gorilla/mux"
)

type Server struct {
	smsHandler  sms.Handler
	ussdHandler ussd.Handler
}

func (s Server) router() *mux.Router {
	r := mux.NewRouter()

	return r
}

func (s Server) start() error {
	srv := http.Server{
		Handler: s.router(),
	}

	if err := srv.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
