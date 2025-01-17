package main

import (
	"log"
	"os"

	_ "github.com/lib/pq" // Библиотека для работы с postgres, driver
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
	todo "github.com/vadimkiryanov/GO-CRUD"
	"github.com/vadimkiryanov/GO-CRUD/pkg/handlers"
	"github.com/vadimkiryanov/GO-CRUD/pkg/repository"
	"github.com/vadimkiryanov/GO-CRUD/pkg/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: [%s]\n", err)
	}

	if err := gotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: [%s]\n", err)
	}

	// Создание подключения к базе данных
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),

		Password: os.Getenv("DB_PASSWORD"), // получение пароля из переменных окружения
	})

	if err != nil {
		log.Fatalf("error initializing db: [%s]\n", err)
	}

	var server = new(todo.Server) // Создание сервера

	var repos = repository.NewRepository(db)     // Создание репозитория
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
