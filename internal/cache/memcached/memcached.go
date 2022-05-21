package memcached

import (
	"context"
	"errors"
)

type Memcached struct {
	Address string
	Port    int
}

func (r Memcached) Get(context.Context, string) (interface{}, error) {
	return nil, errors.New("not implemented")
}

func (r Memcached) Set(context.Context, string, interface{}) error {
	return errors.New("not implemented")
}
