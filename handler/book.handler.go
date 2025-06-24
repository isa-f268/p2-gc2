package handler

import (
	"main/helper"
	"main/model"
	"main/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookHandler struct {
	serv service.BookService
}

func NewBookHandler(s service.BookService) BookHandler {
	return BookHandler{s}
}

func (h *BookHandler) GetBookData(c echo.Context) error {
	idStr := c.QueryParam("genre")

	if idStr == "" {
		book, err := h.serv.GetBookList(nil)
		if err != nil {
			return helper.ErrorHandler(c, 500, err.Error())
		}

		return helper.SuccessHandler(c, 200, "get data book success", book)
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return helper.IdError(c)
	}

	book, err := h.serv.GetBookList(&id)

	if err != nil {
		return helper.ErrorHandler(c, 500, err.Error())
	}

	return helper.SuccessHandler(c, 200, "get data book success", book)
}

func (h *BookHandler) CreateBookLoan(c echo.Context) error {
	var book model.LoanReq

	if err := c.Bind(&book); err != nil {
		return helper.JsonErr(c)
	}

	userId := c.Get("id").(int)

	res, err := h.serv.CreateLoan(book.BookID, userId)

	if err != nil {
		return helper.ErrorHandler(c, 500, err.Error())
	}

	return helper.SuccessHandler(c, 201, "create loan success", res)

}
