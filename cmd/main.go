package main

import (
	"os"

	_ "github.com/lib/pq" // Библиотека для работы с postgres, driver
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
	todo "github.com/vadimkiryanov/GO-CRUD"
	"github.com/vadimkiryanov/GO-CRUD/pkg/handlers"
	"github.com/vadimkiryanov/GO-CRUD/pkg/repository"
	"github.com/vadimkiryanov/GO-CRUD/pkg/service"
)

func main() {
	// Установка уровня логирования
	logrus.SetFormatter(new(logrus.JSONFormatter))

	// Инициализация конфига
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: [%s]\n", err)
	}

	// Инициализация переменных окружения
	if err := gotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: [%s]\n", err)
	}

	// Создание подключения к базе данных
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),     // получение хоста из конфига
		Port:     viper.GetString("db.port"),     // получение порта из конфига
		Username: viper.GetString("db.username"), // получение имени пользователя из конфига
		DBName:   viper.GetString("db.dbname"),   // получение имени базы данных из конфига
		SSLMode:  viper.GetString("db.sslmode"),  // получение режима SSL из конфига

		Password: os.Getenv("DB_PASSWORD"), // получение пароля из переменных окружения
	})

	// Проверка подключения
	if err != nil {
		logrus.Fatalf("error initializing db: [%s]\n", err)
	}

	var server = new(todo.Server) // Создание сервера

	var repos = repository.NewRepository(db)     // Создание репозитория
	var services = service.NewService(repos)     // Создание сервиса
	var handlers = handlers.NewHandler(services) // Создание обработчика

	// Запуск сервера
	// если для viper.GetString key == неверное значение, то запустятся дефолтные настройки
	if err := server.Run(viper.GetString("port"), handlers.InitRouters()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}

// initConfig инициализация конфига
func initConfig() error {
	viper.AddConfigPath("configs") // Папка с конфигами configs/
	viper.SetConfigName("config") // Имя файла с конфигами configs/config.yaml

	return viper.ReadInConfig() // Чтение конфига

}
