package main

import (
	"log"
	"main/config"
	"main/handler"
	"main/middleware"
	"main/repository"
	"main/service"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	godotenv.Load()

	db := config.DBInit()
	e := echo.New()
	//USER
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	authHandler := handler.NewAuthHandler(userService)
	userHandler := handler.NewUserHandller(userService)

	e.POST("/auth/register", authHandler.RegisterUser)
	e.POST("/auth/login", authHandler.LoginUser)

	userGroup := e.Group("/users", middleware.AuthMiddleware)
	userGroup.GET("/detail", userHandler.UserDetail)

	//BOOK
	bookRepo := repository.NewBookRepository(db)
	bookService := service.NewBookRepository(bookRepo)
	bookHandler := handler.NewBookHandler(bookService)

	e.GET("/books", bookHandler.GetBookData, middleware.AuthMiddleware)
	e.POST("/books/loan", bookHandler.CreateBookLoan, middleware.AuthMiddleware)

	//ADMIN
	adminRepo := repository.NewAdminRepository(db)
	adminService := service.NewadminService(adminRepo)
	adminHandler := handler.NewAdminHandler(adminService)

	e.GET("/admin/authors", adminHandler.GetAllAuthor)
	e.GET("/admin/genres", adminHandler.GetGenreDetail)
	e.GET("/admin/users", adminHandler.GetUser)

	log.Println("server running at localhost:8080")
	log.Fatal(e.Start("localhost:8080"))
}
