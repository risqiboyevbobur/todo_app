package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"

	todo "github.com/risqiboyevbobur/todo_app.git"
	"github.com/risqiboyevbobur/todo_app.git/pkg/handler"
	"github.com/risqiboyevbobur/todo_app.git/pkg/repository"
	"github.com/risqiboyevbobur/todo_app.git/pkg/service"
)

func main() {
	if err := initConfigs(); err != nil {
		log.Fatalf("error initilization configs %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
	log.Fatalf("error loading: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	// if err != nil {
	// 	_,err:= repository.NewPostgresDB(repository.Config{
	// 		Host:     "localhost",
	// 		Port:     "5436",
	// 		Username: "postgres",
	// 		Password: "qwerty",
	// 		DBName:   "postgres",
	// 		SSLMode:  "disable",
	// 	})
	// 	// log.Fatalf("faild to initiliazing db %s", err.Error())
	fmt.Printf(err.Error())
	// }
	repos := repository.NewRepository(db)
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
