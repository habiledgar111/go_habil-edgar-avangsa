package main

import (
	"sec_23/Praktikum/restAPI_unit_testing/controller"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.POST("/users/login", controller.UserLogin)
	e.GET("/sayhello", controller.SayHello)
	e.GET("/users", controller.GetAllUser, echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte("secret"),
	}))
	e.POST("/users", controller.CreateUser)
	e.PUT("/users/:id", controller.UpdateUser, echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte("secret"),
	}))
	e.DELETE("users/:id", controller.DeleteUser, echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte("secret"),
	}))
	// e.Use(echojwt.WithConfig(echojwt.Config{
	// 	SigningKey: []byte("secret"),
	// }))
	e.GET("/users/:id", controller.GetUser, echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte("secret"),
	}))
	e.Logger.Fatal(e.Start(":1323"))
}
