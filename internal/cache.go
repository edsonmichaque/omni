package internal

import "context"

type Cache interface {
	Get(context.Context, string) (interface{}, error)
	Set(context.Context, string, interface{}) error
}
