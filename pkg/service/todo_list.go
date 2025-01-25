package service

import (
	todo "github.com/vadimkiryanov/GO-CRUD"
	"github.com/vadimkiryanov/GO-CRUD/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (service *TodoListService) Create(userId int, list todo.TodoList) (int, error) {
	return service.repo.Create(userId, list)
}
func (service *TodoListService) GetAll(userId int) ([]todo.TodoList, error) {
	return service.repo.GetAll(userId)
}

func (service *TodoListService) GetById(userId, listId int) (todo.TodoList, error) {
	return service.repo.GetById(userId, listId)
}
