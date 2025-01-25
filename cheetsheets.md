[3]
- go get -u github.com/spf13/viper // библиотека для работы с конфигами
- docker pull postgres // Скачивание образа postgres из докер хаба для работы с postgres

- docker run --name=todo-db -e POSTGRES_PASSWORD=qwerty123 -p 5432:5432 --rm postgres // Создание контейнера с его удалением после остановки
- docker run --name=todo-db -e POSTGRES_PASSWORD=qwerty123 -p 5432:5432 -v ${HOME}/pgdata/:/var/lib/postgresql/data postgres // Персист данных и сохранение контейнера после остановки 

[
    docker run - базовая команда для запуска нового контейнера

    Далее идут параметры:

    --name=todo-db - задает имя контейнеру ("todo-db")
    -e POSTGRES_PASSWORD=qwerty123 - устанавливает переменную окружения для пароля PostgreSQL
    -p 5432:5432 - проброс портов, где:
    первый 5432 - порт на хост-машине
    второй 5432 - порт внутри контейнера (это позволяет подключаться к базе данных с хост-машины через порт 5432)
    --rm - автоматически удаляет контейнер после его остановки
    postgres - имя образа PostgreSQL, который будет использоваться
]

- go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest // Установка через golang утилиты "migrate"
- migrate create -ext sql -dir ./schema -seq init
[
    migrate create - создает новый файл миграции

    Параметры:
    -ext sql - указывает расширение файла (в данном случае .sql)
    -dir ./schema - указывает директорию, где будут созданы файлы миграции
    -seq init - задает название миграции ("init") и использует последовательную нумерацию
    В результате будут созданы два файла в директории ./schema:

    NNNNNN_init.up.sql - файл для применения миграции (внесения изменений)
    NNNNNN_init.down.sql - файл для отката миграции (отмены изменений)
    где NNNNNN - это порядковый номер миграции (например, 000001)
]

- migrate -path ./schema -database 'postgres://postgres:qwerty123@localhost:5432/postgres?sslmode=disable' up
[
    migrate - это команда инструмента для управления миграциями
    -path ./shema - указывает путь к директории, где хранятся файлы миграций (в данном случае это папка "shema")
    -database - флаг для указания строки подключения к базе данных
    'postgres://postgres:qwerty123@localhost:5436/postgres?sslmode=disable' - это строка подключения к PostgreSQL базе данных, где:
    postgres:// - протокол подключения
    postgres - имя пользователя
    qwerty123 - пароль
    localhost - хост (локальный компьютер)
    5436 - порт
    postgres - имя базы данных
    sslmode=disable - отключение SSL-соединения
    up - команда для применения всех доступных миграций, которые еще не были применены
]

- docker exec -it [CONTAINER_ID] /bin/bash // подключение к бд
- psql -U postgres // подключение к бд
- \d // Просмотр всех таблиц
- exit // выход с бд

- migrate -path ./schema -database 'postgres://postgres:qwerty123@localhost:5432/postgres?sslmode=disable' down // 
- update schema_migrations set version='000001', dirty=false; // исправление dirty бд

[5] Подключение к БД из приложения. Переменные окружения. Библиотека sqlx
go get -u github.com/jmoiron/sqlx // 
go get -u github.com/joho/godotenv // Для работы с env переменными 

[6] Регистрация пользователей
- go get -u github.com/sirupsen/logrus // Библиотека для работы с логами 

[7] Аутентификация. JWT-токены.
- go get -u github.com/golang-jwt/jwt // Библиотека для работы с JWT 

[9] - Транзакции - это выполнение нескольких операций, (вставка, вставка в таблицы)