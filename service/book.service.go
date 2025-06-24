package service

import (
	"main/model"
	"main/repository"
)

type BookService interface {
	GetBookList(id *int) ([]model.Book, error)
	CreateLoan(bookID int, userID int) (model.Book, error)
}

type bookService struct {
	repo repository.BookRepository
}

func NewBookRepository(r repository.BookRepository) BookService {
	return &bookService{r}
}

func (s *bookService) GetBookList(id *int) ([]model.Book, error) {
	return s.repo.GetBook(id)
}

func (s *bookService) CreateLoan(bookID int, userID int) (model.Book, error) {
	return s.repo.CreateLoan(bookID, userID)
}
