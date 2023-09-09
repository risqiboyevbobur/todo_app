package main

import (
	"log"

	todo "github.com/risqiboyevbobur/todo_app.git"
	"github.com/risqiboyevbobur/todo_app.git/pkg/handler"
)

func main()  {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000",handlers.InitRoutes()); err != nil {
		log.Fatalf("error accured while running http server: %s", err.Error())
	}
}