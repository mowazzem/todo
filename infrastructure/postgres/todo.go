package postgres

import (
	"context"
	"log"
	"time"

	"github.com/mowazzem/lets-go/domain"
	"gorm.io/gorm"
)

type todoRepoImpl struct {
	DB *gorm.DB
}

func NewTodoRepoImpl(db *gorm.DB) domain.TodoRepo {
	tdimpl := todoRepoImpl{
		DB: db,
	}

	return tdimpl
}

type todo struct {
	ID          string
	Name        string
	Description string
	CreatedTime time.Time
	UpdatedTime time.Time
	DeletedTime *time.Time
}

func (todo) TableName() string {
	return "todo"
}

func toTodoDto(do domain.Todo) todo {
	return todo{
		ID:          do.ID.String(),
		Name:        do.Name,
		Description: do.Description,
	}
}

func (repo todoRepoImpl) Create(ctx context.Context, todo domain.Todo) (*domain.Todo, error) {
	td := toTodoDto(todo)

	log.Println(todo)

	err := repo.DB.Create(&td).Error
	if err != nil {
		return nil, err
	}

	return &todo, nil
}
