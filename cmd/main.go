package main

import (
	todo "github.com/m0rk0vka/TodoApp"
	"github.com/m0rk0vka/TodoApp/pkg/handler"
	"github.com/m0rk0vka/TodoApp/pkg/repository"
	"github.com/m0rk0vka/TodoApp/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

}
