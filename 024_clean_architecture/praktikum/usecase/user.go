package usecase

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sec024/praktikum/entity"
	"sec024/praktikum/repository"

	"github.com/labstack/echo/v4"
)

func GetAll(c echo.Context) error {
	user, err := repository.GetAllUser()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "server problem")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
}

func Create(c echo.Context) error {
	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Massage": "json cant empty",
		})
	}
	email := fmt.Sprintf("%v", json_map["email"])
	pass := fmt.Sprintf("%v", json_map["pass"])
	user := entity.User{
		Email:    email,
		Password: pass,
	}

	err = repository.Createuser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "cant create users",
		})
	}
	return c.JSON(http.StatusOK, "sucess")
}
