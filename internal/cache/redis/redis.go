package redis

import (
	"context"
	"errors"
)

type Redis struct {
	Address string
	Port    int
}

func (r Redis) Get(context.Context, string) (interface{}, error) {
	return nil, errors.New("not implemented")
}

func (r Redis) Set(context.Context, string, interface{}) error {
	return errors.New("not implemented")
}
