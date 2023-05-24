package main

import (
	"code_competence/config"
	"code_competence/controller"
	"code_competence/model"
	"net/http"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	barang   *model.Barang
	kategori *model.Kategori
	user     *model.User
)

func main() {

	config.Open()
	config.DB.AutoMigrate(user, kategori)
	config.DB.AutoMigrate(barang)
	// config.DB.AutoMigrate(&model.Barang{})
	// config.DB.AutoMigrate(&model.User{})

	e := echo.New()
	e.Use(middleware.Logger())

	var middlewareJWT = echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte("secret"),
	})

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/barang", controller.GetAllBarang, middlewareJWT)
	e.GET("/barang/:id", controller.GetBarangID, middlewareJWT)
	e.POST("/barang", controller.CreateBarang, middlewareJWT)
	e.PUT("/barang/:id", controller.UpdateBarang, middlewareJWT)
	e.DELETE("/barang/:id", controller.DeleteBarang, middlewareJWT)
	e.GET("/barangname", controller.GetBarangfromName, middlewareJWT)

	e.GET("/items/category/:id", controller.GetAllBarangfromKategori, middlewareJWT)

	e.POST("/user/login", controller.Login)
	e.POST("/user/register", controller.CreateUser)
	e.Logger.Fatal(e.Start(":1323"))
}
