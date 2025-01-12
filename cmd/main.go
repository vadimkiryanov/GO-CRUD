package main

import (
	"log"

	todo "github.com/vadimkiryanov/GO-CRUD"
	"github.com/vadimkiryanov/GO-CRUD/pkg/handlers"
)

func main() {
	var server = new(todo.Server)
	var handlers = new(handlers.Handler)

	if err := server.Run("8000", handlers.InitRouters()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
