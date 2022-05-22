package internal

import "context"

type SecretRetriever interface {
	Retrieve(context.Context, string) (string, error)
}

type Secret interface {
	SecretRetriever
}
