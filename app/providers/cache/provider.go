package cache

import (
	"context"
	"time"
)

type Provider struct {
	client interface{}
	ctx    context.Context
}

type IProvider interface {
	Set(key string, value interface{}, exp time.Duration) error
	Get(key string) (interface{}, error)
}
