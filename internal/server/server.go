package server

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/edsonmichaque/omni/internal"
	"github.com/edsonmichaque/omni/internal/cache/memcached"
	"github.com/edsonmichaque/omni/internal/cache/redis"
	"github.com/edsonmichaque/omni/internal/config"
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

func (s Server) start(c *config.Config) error {
	cache, err := newCache(c.Cache)
	if err != nil {
		return err
	}

	logger, err := newLogger(c.Logger)
	if err != nil {
		return err
	}

	s.smsHandler = sms.Handler{
		Cache:  cache,
		Logger: logger,
	}

	s.ussdHandler = ussd.Handler{
		Cache:  cache,
		Logger: logger,
	}

	srv := http.Server{
		Addr:    fmt.Sprintf("%s:%d", c.Address, c.Port),
		Handler: s.router(),
	}

	if err := srv.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func newCache(c config.Cache) (internal.Cache, error) {
	if c.Redis != (config.Redis{}) {
		return redis.Redis{}, nil
	}

	if c.Memcached != (config.Memcached{}) {
		return memcached.Memcached{}, nil
	}

	return nil, errors.New("invalid cache")
}

func newLogger(c config.Logger) (internal.Logger, error) {
	return nil, errors.New("invalid logger")
}
