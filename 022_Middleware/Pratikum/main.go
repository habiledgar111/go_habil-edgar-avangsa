package main

import (
	"sec_22/022_Middleware/Pratikum/controller"
	midd "sec_22/022_Middleware/Pratikum/middleware"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}",` +
			`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
			`"status":${status},"error":"${error}","latency":${latency},"latency_human":"${latency_human}"` +
			`,"bytes_in":${bytes_in},"bytes_out":${bytes_out}}` + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
	}))
	e.GET("/sayhello", controller.SayHello)
	e.POST("/user/login", controller.UserLogin)
	e.GET("/users", controller.GetAllUser)
	e.GET("/users/:id", controller.GetUser)
	e.POST("/users", controller.CreateUser)
	e.PUT("/users/:id", controller.UpdateUser)
	e.DELETE("users/:id", controller.DeleteUser)
	e.GET("/books", controller.GetAllBooks)
	e.POST("/books", controller.CreateBook)
	e.PUT("/books/:id", controller.UpdateBook)
	e.DELETE("books/:id", controller.DeleteBook)
	// e.Use(middleware.echojwt.WithConfig(echojwt.Config{
	// 	SigningKey: []byte("secret"),
	// }))
	// e.Use(midd.ExtractJWT)
	e.GET("/books/:id", controller.GetBook, midd.ExtractJWT)
	e.Logger.Fatal(e.Start(":1323"))
}
