package todo

import (
	"context"

	"github.com/mowazzem/lets-go/domain"
)

func (uc usecase) Create(ctx context.Context, todo domain.Todo) (*domain.Todo, error) {
	return uc.todoRepo.Create(ctx, todo)
}
