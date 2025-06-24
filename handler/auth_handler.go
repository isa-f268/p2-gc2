package handler

import (
	"main/helper"
	"main/model"
	"main/service"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	serv service.UserService
}

func NewAuthHandler(s service.UserService) *AuthHandler {
	return &AuthHandler{s}
}

// POST /users/register

func (h *AuthHandler) RegisterUser(c echo.Context) error {
	var u model.User

	if err := c.Bind(&u); err != nil {
		return helper.InvalidReq(c)
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	u.Password = string(hashed)
	user, err := h.serv.UserCreate(u)
	if err != nil {
		return helper.InternalErr(c)
	}

	res := model.RespRegister{
		User_id: int(user.ID),
		Name:    user.FirstName,
	}

	return c.JSON(201, res)
}

// POST /users/login
func (h *AuthHandler) LoginUser(c echo.Context) error {
	var u model.LoginReq

	if err := c.Bind(&u); err != nil {
		return helper.InvalidReq(c)
	}
	resp, err := h.serv.UserLogin(u)
	if err != nil {
		return helper.InternalErr(c)
	}

	token := map[string]string{"token": resp}

	return c.JSON(200, token)
}
