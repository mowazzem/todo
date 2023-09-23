package http

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/mowazzem/lets-go/domain"
	todoUc "github.com/mowazzem/lets-go/usecase/todo"
)

type todoDeliver struct {
	router chi.Router
	uc     todoUc.Usecase
}

func NewTodoDeliver(router chi.Router, uc todoUc.Usecase) {
	tdDeliver := todoDeliver{
		router: router,
		uc:     uc,
	}
	router.Post("/create-todo", tdDeliver.Create)
}

type todo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type restodo struct {
	ID string `json:"id"`
}

func (tdo todo) toDomain() (*domain.Todo, error) {
	dtdo, err := domain.NewTodo(tdo.ID, tdo.Name, tdo.Desc)
	if err != nil {
		return nil, err
	}

	return &dtdo, nil
}

func (td todoDeliver) Create(w http.ResponseWriter, r *http.Request) {
	var tdo todo
	json.NewDecoder(r.Body).Decode(&tdo)

	dtdo, err := tdo.toDomain()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	created, err := td.uc.Create(ctx, *dtdo)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	res := restodo{
		ID: created.ID.String(),
	}
	json.NewEncoder(w).Encode(res)
}
