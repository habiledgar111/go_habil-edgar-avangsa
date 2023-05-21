package main

import (
	"code_competence/controller"
	"net/http"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	var middlewareJWT = echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte("secret"),
	})

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/barang", controller.GetAllBarang, middlewareJWT)
	e.GET("/barang/:id", controller.GetBarangID, middlewareJWT)
	e.POST("/barang", controller.CreateBarang, middlewareJWT)
	e.DELETE("/barang/:id", controller.DeleteBarang, middlewareJWT)

	e.POST("/user/login", controller.Login)
	e.POST("/user/register", controller.CreateUser)
	e.Logger.Fatal(e.Start(":1323"))
}
