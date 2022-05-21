package server

import (
	"fmt"
	"net/http"

	"github.com/edsonmichaque/omni/internal"
	"github.com/edsonmichaque/omni/internal/cache/memcached"
	"github.com/edsonmichaque/omni/internal/cache/redis"
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

func (s Server) start(c *internal.Config) error {
	var cache internal.Cache

	if c.Cache.Redis != (internal.RedisConfig{}) {
		cache = redis.Redis{}
	}

	if c.Cache.Memcached != (internal.MemcachedConfig{}) {
		cache = memcached.Memcached{}
	}

	s.smsHandler = sms.Handler{
		Cache: cache,
	}

	s.ussdHandler = ussd.Handler{
		Cache: cache,
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
