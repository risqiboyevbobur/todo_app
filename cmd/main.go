package main

import (
	"log"

	"github.com/risqiboyevbobur/todo_app.git"
	"github.com/risqiboyevbobur/todo_app.git/pkg/handler"
	"github.com/risqiboyevbobur/todo_app.git/pkg/repository"
	"github.com/risqiboyevbobur/todo_app.git/pkg/service"
)

func main()  {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(todo.Server)
	if err := srv.Run("8000",handlers.InitRoutes()); err != nil {
		log.Fatalf("error accured while running http server: %s", err.Error())
	}
}