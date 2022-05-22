package server

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/edsonmichaque/omni/internal"
	"github.com/edsonmichaque/omni/internal/cache/memcached"
	"github.com/edsonmichaque/omni/internal/cache/redis"
	"github.com/edsonmichaque/omni/internal/config"
	"github.com/edsonmichaque/omni/internal/server/message"
	"github.com/edsonmichaque/omni/internal/server/session"
	"github.com/gorilla/mux"
)

type Server struct {
	smsHandler  message.Handler
	ussdHandler session.Handler
}

func (s Server) router() *mux.Router {
	r := mux.NewRouter()

	return r
}

func (s Server) start(c config.Server) error {
	cache, err := newCacheProvider(c)
	if err != nil {
		return err
	}

	logger, err := newLogProvider(c)
	if err != nil {
		return err
	}

	s.smsHandler = message.Handler{
		Cache:  cache,
		Logger: logger,
	}

	s.ussdHandler = session.Handler{
		Cache:  cache,
		Logger: logger,
	}

	srv := http.Server{
		Addr:    fmt.Sprintf("%s:%d", c.Server.Address, c.Server.Port),
		Handler: s.router(),
	}

	if err := srv.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func newCacheProvider(c config.Server) (internal.Cache, error) {
	if c.Cache.Redis != nil {
		return redis.Redis{}, nil
	}

	if c.Cache.Memcached != nil {
		return memcached.Memcached{}, nil
	}

	return nil, errors.New("invalid cache")
}

func newLogProvider(c config.Server) (internal.Log, error) {
	return nil, errors.New("invalid logger")
}
