package service

import (
	"github.com/zhashkevych/todo-app/models"
	"github.com/zhashkevych/todo-app/package/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Update(userId int, input models.UpdateUser) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(userId, input)
}

func (s *UserService) Get(userId int) (models.User, error) {
	return s.repo.Get(userId)
}

func (s *UserService) Delete(userId int) error {
	return s.repo.Delete(userId)
}
