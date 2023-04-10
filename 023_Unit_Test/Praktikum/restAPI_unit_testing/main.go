package main

import (
	"sec_23/Praktikum/restAPI_unit_testing/controller"

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
	e.Logger.Fatal(e.Start(":1323"))
}
