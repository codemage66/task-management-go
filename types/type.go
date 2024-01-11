package types

import (
	"context"
)

type Repository interface {
	Get(context.Context, interface{}) (interface{}, error)
	List(context.Context, Pageable) (interface{}, interface{}, error)
	Create(context.Context, interface{}) (interface{}, error)
	Update(context.Context, interface{}, interface{}) (interface{}, error)
	Delete(context.Context, interface{}) error
}

type Env struct {
	Port    string
	Host    string
	DBUrl   string
	SSLMode string
}
