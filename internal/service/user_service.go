package service

import (
	"myapp/internal/model"
	"myapp/internal/repository"
)

func CreateUser(user model.User) error {
	return repository.CreateUser(user)
}
