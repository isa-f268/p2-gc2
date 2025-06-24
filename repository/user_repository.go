package repository

import (
	"fmt"
	"main/helper"
	"main/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user model.User) (model.User, error)
	LoginUser(login model.LoginReq) (string, error)
	GetUserDetail(id int) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateUser(u model.User) (model.User, error) {
	user := model.User{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Address:   u.Address,
		DateBirth: u.DateBirth,
		Email:     u.Email,
		Password:  u.Password,
	}
	res := r.db.Create(&user)

	if res.Error != nil {
		return model.User{}, fmt.Errorf("fail to create new user")
	}

	return user, nil
}

func (r *userRepository) LoginUser(login model.LoginReq) (string, error) {
	var u model.User

	res := r.db.Where("email=?", login.Email).First(&u)

	if res.Error != nil {
		return "", fmt.Errorf("invalid email")
	}

	err := helper.CheckPassword(u.Password, login.Password)

	if err != nil {
		return "", fmt.Errorf("password error")
	}

	data := helper.Login{
		Id:   int(u.ID),
		Name: u.FirstName,
	}

	token, err := helper.CreateJwt(data)

	if err != nil {
		return "", fmt.Errorf("fail to provide data")
	}

	return token, nil
}

func (r *userRepository) GetUserDetail(id int) (model.User, error) {
	var u model.User

	res := r.db.Preload("Loan.Book.Genre").Where("id=?", id).First(&u)

	if res.Error != nil {
		return model.User{}, fmt.Errorf("user not found")
	}

	return u, nil
}
