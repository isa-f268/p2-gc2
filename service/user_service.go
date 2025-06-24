package service

import (
	"main/model"
	"main/repository"
)

type UserService interface {
	UserCreate(u model.User) (model.User, error)
	UserLogin(u model.LoginReq) (string, error)
	UserDetail(id int) (model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{r}
}

func (s *userService) UserCreate(u model.User) (model.User, error) {
	return s.repo.CreateUser(u)
}

func (s *userService) UserLogin(u model.LoginReq) (string, error) {
	return s.repo.LoginUser(u)
}

func (s *userService) UserDetail(id int) (model.User, error) {
	return s.repo.GetUserDetail(id)
}
