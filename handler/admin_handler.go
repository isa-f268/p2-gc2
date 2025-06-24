package handler

import (
	"main/helper"
	"main/service"

	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	serv service.AdminServie
}

func NewAdminHandler(s service.AdminServie) *AdminHandler {
	return &AdminHandler{s}
}

// GET /admin/authors
func (h *AdminHandler) GetAllAuthor(c echo.Context) error {
	authors, err := h.serv.Getauthor()
	if err != nil {
		return helper.ErrorHandler(c, 500, err.Error())
	}

	return helper.SuccessHandler(c, 200, "get author data success", authors)

}

// GET /admin/genres
func (h *AdminHandler) GetGenreDetail(c echo.Context) error {
	genres, err := h.serv.GenreDetail()
	if err != nil {
		return helper.ErrorHandler(c, 500, err.Error())
	}

	return helper.SuccessHandler(c, 200, "get genre detail success", genres)
}

// GET /admin/users
func (h *AdminHandler) GetUser(c echo.Context) error {
	users, err := h.serv.UserDetail()
	if err != nil {
		return helper.ErrorHandler(c, 500, err.Error())
	}

	return helper.SuccessHandler(c, 200, "get genre detail success", users)

}
