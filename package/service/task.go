package service

import (
	"github.com/zhashkevych/todo-app/models"
	"github.com/zhashkevych/todo-app/package/repository"
)

type TaskService struct {
	repo repository.Task
}

func NewTaskService(repo repository.Task) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) Create(userId int, task models.TaskCreateInput) (int, error) {
	return s.repo.Create(userId, task)
}

func (s *TaskService) GetAll(userId int) ([]models.Task, error) {
	return s.repo.GetAll(userId)
}

func (s *TaskService) GetById(userId, taskId int) (models.Task, error) {
	return s.repo.GetById(userId, taskId)
}

func (s *TaskService) Update(userId, taskId int, input models.TaskUpdateInput) error {
	return s.repo.Update(userId, taskId, input)
}

func (s *TaskService) Delete(userId, taskId int) error {
	return s.repo.Delete(userId, taskId)
}
