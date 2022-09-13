package main

import (
	"log"
	"todo"
	"todo/pkg/handler"
	"todo/pkg/repository"
	"todo/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	svr := new(todo.Server)
	if err := svr.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while ruuning http server: %s", err.Error())
	}
}
