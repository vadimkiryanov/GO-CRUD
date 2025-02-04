package repository

import (
	"github.com/jmoiron/sqlx"
	todo "github.com/vadimkiryanov/GO-CRUD"
)

// Authorization интерфейс репозитория для работы с авторизацией
type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId, listId int) (todo.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input todo.UpdateListInput) error
}

type TodoItem interface {
}
type Repository struct {
	Authorization
	TodoList
	TodoItem
}

// NewRepository создает новый экземпляр структуры Repository
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
	}
}
