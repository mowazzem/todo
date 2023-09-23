package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	todoHandler "github.com/mowazzem/lets-go/deliver/http"
	"github.com/mowazzem/lets-go/infrastructure/postgres"
	"github.com/mowazzem/lets-go/usecase/todo"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	db, err := postgres.Connect()
	if err != nil {
		panic(err)
	}

	todoRepo := postgres.NewTodoRepoImpl(db)
	todoUc := todo.NewUsecase(todoRepo)
	todoHandler.NewTodoDeliver(r, todoUc)

	log.Println("server listening at:3000")
	http.ListenAndServe(":3000", r)
}
