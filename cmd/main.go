package main

import (
	"log"

	"github.com/spf13/viper"

	todo "github.com/risqiboyevbobur/todo_app.git"
	"github.com/risqiboyevbobur/todo_app.git/pkg/handler"
	"github.com/risqiboyevbobur/todo_app.git/pkg/repository"
	"github.com/risqiboyevbobur/todo_app.git/pkg/service"
)

func main() {
	if err := initConfigs(); err != nil{
		log.Fatalf("error initilization configs %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("8080"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error accured while running http server: %s", err.Error())
	}
}
func initConfigs() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
