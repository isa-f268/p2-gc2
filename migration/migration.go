package main

import (
	"main/config"
	"main/model"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db := config.DBInit()

	db.AutoMigrate(
		&model.User{},
		&model.Author{},
		&model.Genre{},
		&model.Book{},
		&model.Loan{},
	)

}
