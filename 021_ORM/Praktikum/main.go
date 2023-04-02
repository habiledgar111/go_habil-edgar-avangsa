package main

import (
	"021_ORM/021_ORM/Praktikum/controller"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/sayhello", controller.SayHello)
	e.Logger.Fatal(e.Start(":1323"))
}
