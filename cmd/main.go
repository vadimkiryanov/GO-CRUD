package main

import (
	"log"

	"github.com/spf13/viper"
	todo "github.com/vadimkiryanov/GO-CRUD"
	"github.com/vadimkiryanov/GO-CRUD/pkg/handlers"
	"github.com/vadimkiryanov/GO-CRUD/pkg/repository"
	"github.com/vadimkiryanov/GO-CRUD/pkg/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: [%s]\n", err)
	}

	var server = new(todo.Server) // Создание сервера

	var repos = repository.NewRepository()       // Создание репозитория
	var services = service.NewService(repos)     // Создание сервиса
	var handlers = handlers.NewHandler(services) // Создание обработчика

	// Запуск сервера
	// если для viper.GetString key == неверное значение, то запустятся дефолтные настройки
	if err := server.Run(viper.GetString("port"), handlers.InitRouters()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()

}
