package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `gorm:"not null" json:"first_name"`
	LastName  string `gorm:"not null" json:"last_name"`
	Address   string `gorm:"not null" json:"address"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	DateBirth string `gorm:"type:date" json:"birth"`
	Loan      []Loan `gorm:"foreignKey:UserID"`
}

type Author struct {
	gorm.Model
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Book      []Book `gorm:"foreignKey:AuthorID"`
}

type Genre struct {
	gorm.Model
	GenreName string `gorm:"not null"`
	Book      []Book `gorm:"foreignKey:GenreID"`
}

type Book struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	MinAge      string `gorm:"not null"`
	CoverUrl    string `gorm:"not null"`
	Loan        []Loan `gorm:"foreignKey:BookID"`
	AuthorID    uint
	GenreID     uint

	Genre Genre
}

type Loan struct {
	gorm.Model
	UserID uint
	BookID uint

	Book Book
}
