package domain

import "context"

type TodoRepo interface {
	Create(ctx context.Context, todo Todo) (*Todo, error)
}
