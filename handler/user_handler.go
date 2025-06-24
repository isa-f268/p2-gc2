package handler

import (
	"main/helper"
	"main/service"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	serv service.UserService
}

func NewUserHandller(s service.UserService) *UserHandler {
	return &UserHandler{s}
}

// GET /users/detail
func (h *UserHandler) UserDetail(c echo.Context) error {
	id := c.Get("id").(int)
	res, err := h.serv.UserDetail(id)

	if err != nil {
		return helper.IdError(c)
	}

	return helper.SuccessHandler(c, 200, "get user success", res)
}
