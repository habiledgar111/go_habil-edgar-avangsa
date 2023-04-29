package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

var (
	// lock = sync.Mutex{}

	//tempat nyimpan data
	users = make(map[int]user)
)

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func createUser(ctx echo.Context) error {
	// lock.Lock()
	// defer lock.Unlock()
	json_map := make(map[string]interface{})
	err := json.NewDecoder(ctx.Request().Body).Decode(&json_map)
	if err != nil {
		log.Error("Json is empty")
		massage := struct {
			Massage string `json:"massage"`
		}{
			Massage: "Json cant empty",
		}
		return ctx.JSON(http.StatusBadRequest, massage)
	}
	name := fmt.Sprintf("%v", json_map["name"])
	age, _ := strconv.Atoi(fmt.Sprintf("%v", json_map["age"]))
	u := user{
		ID:   len(users) + 1,
		Name: name,
		Age:  age,
	}
	massage := struct {
		Massage string `json:"massage"`
		User    user   `json:"data"`
	}{
		Massage: "berhasil menambahkan data",
		User:    u,
	}
	users[u.ID] = u
	return ctx.JSON(http.StatusOK, massage)
}

func getAllUsers(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "return all user",
		"users":   users,
	})
}

func deleteUser(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	for key, _ := range users {
		if key == id {
			delete(users, id)
			return ctx.JSON(http.StatusOK, map[string]interface{}{
				"message": "sucsses delete data",
			})
		}
	}
	return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
		"message": "cant find id",
	})
}

func updateUser(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	for key, _ := range users {
		if key == id {
			json_map := make(map[string]interface{})
			err := json.NewDecoder(ctx.Request().Body).Decode(&json_map)
			if err != nil {
				log.Error("Json is empty")
				massage := struct {
					Massage string `json:"massage"`
				}{
					Massage: "Json cant empty",
				}
				return ctx.JSON(http.StatusBadRequest, massage)
			}
			name := fmt.Sprintf("%v", json_map["name"])
			age, _ := strconv.Atoi(fmt.Sprintf("%v", json_map["age"]))
			u := user{
				ID:   len(users) + 1,
				Name: name,
				Age:  age,
			}
			users[id] = u
			return ctx.JSON(http.StatusOK, map[string]interface{}{
				"message": "sucsses update data",
			})
		}
	}
	return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
		"message": "cant find id",
	})
}
func getUserID(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	for key, _ := range users {
		if key == id {
			return ctx.JSON(http.StatusOK, users[id])
		}
	}
	return ctx.JSON(http.StatusBadRequest, map[string]string{
		"massage": "cant find id",
	})
}
func main() {
	e := echo.New()
	e.POST("/users", createUser)
	e.GET("/users", getAllUsers)
	e.GET("/users/:id", getUserID)
	e.DELETE("/users/:id", deleteUser)
	e.PUT("/users/:id", updateUser)
	fmt.Println("listen to request")
	e.Logger.Fatal(e.Start(":1323"))
}
