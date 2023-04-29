package main

import (
	"021_ORM/021_ORM/Praktikum/controller"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/sayhello", controller.SayHello)
	e.GET("/users", controller.GetAllUser)
	e.GET("/users/:id", controller.GetUser)
	e.POST("/users", controller.CreateUser)
	e.PUT("/users/:id", controller.UpdateUser)
	e.DELETE("users/:id", controller.DeleteUser)
	e.GET("/books", controller.GetAllBooks)
	e.GET("/books/:id", controller.GetBook)
	e.POST("/books", controller.CreateBook)
	e.PUT("/books/:id", controller.UpdateBook)
	e.DELETE("books/:id", controller.DeleteBook)
	e.Logger.Fatal(e.Start(":1323"))
}
