package main

import (
	"sec024/praktikum/config"
	"sec024/praktikum/usecase"

	"github.com/labstack/echo/v4"
)

func main() {
	config.Connection()

	e := echo.New()
	e.GET("/users", usecase.GetAll)
	e.POST("/users", usecase.Create)
	e.Start(":1323")
}
