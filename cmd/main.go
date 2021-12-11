package main

import (
	todo "github.com/KrissLa/todo_app_go"
	"github.com/KrissLa/todo_app_go/pkg/handler"
	"github.com/KrissLa/todo_app_go/pkg/repository"
	"github.com/KrissLa/todo_app_go/pkg/service"
	"log"
)

func main()  {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatal("error occured while running http server: %s", err.Error())
	}
}
