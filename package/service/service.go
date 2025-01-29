package service

import (
	"github.com/zhashkevych/todo-app/models"
	"github.com/zhashkevych/todo-app/package/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list models.TodoList) (int, error)
	GetAll(userId int) ([]models.TodoList, error)
	GetById(userId, listId int) (models.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input models.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId int, list models.TodoItem) (int, error)
	GetAll(userId, listId int) ([]models.TodoItem, error)
	GetById(userId, itemId int) (models.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input models.UpdateItemInput) error
}

type User interface {
	Update(userId int, input models.UpdateUser) error
	Get(userId int) (models.User, error)
	Delete(userId int) error
}

type Task interface {
	Create(userId int, task models.TaskCreateInput) (int, error)
	GetAll(userId int) ([]models.Task, error)
	GetById(userId, taskId int) (models.Task, error)
	Update(userId, taskId int, input models.TaskUpdateInput) error
	Delete(userId, taskId int) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
	User
	Task
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
		User:          NewUserService(repos.User),
		Task:          NewTaskService(repos.Task),
	}
}
