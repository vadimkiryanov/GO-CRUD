package main

import (
	"log"

	todo "github.com/vadimkiryanov/GO-CRUD"
	"github.com/vadimkiryanov/GO-CRUD/pkg/handlers"
	"github.com/vadimkiryanov/GO-CRUD/pkg/repository"
	"github.com/vadimkiryanov/GO-CRUD/pkg/service"
)

func main() {
	var server = new(todo.Server)

	var repos = repository.NewRepository()
	var services = service.NewService(repos)
	var handlers = handlers.NewHandler(services)

	if err := server.Run("8000", handlers.InitRouters()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
