package todo

import (
	"context"

	"github.com/mowazzem/lets-go/domain"
)

type Usecase interface {
	Create(ctx context.Context, todo domain.Todo) (*domain.Todo, error)
}

type usecase struct {
	todoRepo domain.TodoRepo
}

func NewUsecase(tr domain.TodoRepo) Usecase {
	return usecase{
		todoRepo: tr,
	}
}
