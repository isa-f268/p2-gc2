package repository

import (
	"fmt"
	"main/model"

	"gorm.io/gorm"
)

type BookRepository interface {
	GetBook(id *int) ([]model.Book, error)
	CreateLoan(bookID int, userID int) (model.Book, error)
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db}
}

func (r bookRepository) GetBook(id *int) ([]model.Book, error) {
	var books []model.Book

	if id != nil {
		res := r.db.Where("id=?", id).Find(&books)

		if res.Error != nil {
			return []model.Book{}, fmt.Errorf("get data error")
		}

		return books, nil
	}

	res := r.db.Find(&books)
	if res.Error != nil {
		return []model.Book{}, fmt.Errorf("get data error")
	}

	return books, nil
}

func (r bookRepository) CreateLoan(bookID int, userID int) (model.Book, error) {
	var book model.Book
	loan := model.Loan{
		UserID: uint(userID),
		BookID: uint(bookID),
	}

	res := r.db.Create(&loan)
	r.db.Where("id=?", bookID).First(&book)

	if res.Error != nil {
		return model.Book{}, fmt.Errorf("fail to create data")
	}

	return book, nil

}
