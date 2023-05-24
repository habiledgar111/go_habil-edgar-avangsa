package controller

import (
	"code_competence/config"
	"code_competence/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/labstack/echo/v4"
)

type jwtcustomclaims struct {
	Email    string
	Password string
	jwt.RegisteredClaims
}

func createJWT(user model.User) interface{} {
	email := user.Email
	pass := user.Password
	claims := &jwtcustomclaims{
		email,
		pass,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	temp := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := temp.SignedString([]byte("secret"))

	if err != nil {
		return err.Error()
	}

	return token
}

func Login(c echo.Context) error {
	var users model.User
	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err,
		})
	}

	email := fmt.Sprintf("%v", json_map["email"])
	password := fmt.Sprintf("%v", json_map["password"])

	// user = model.GetUserFromEmail(email)
	if err := config.DB.Where("email = @email", sql.Named("email", email)).First(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if users.Password != password {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "wrong password",
		})
	}
	token := createJWT(users)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"token":   token,
	})
}

func CreateUser(c echo.Context) error {
	var users model.User
	json_map := make(map[string]interface{})

	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Massage": "json cant empty",
		})
	}

	email := fmt.Sprintf("%v", json_map["email"])
	password := fmt.Sprintf("%v", json_map["password"])

	users = model.User{
		Email:    email,
		Password: password,
	}

	result := config.DB.Create(&users)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":     "success create user",
		"rowaffected": result.RowsAffected,
		"querybool":   result.QueryFields,
	})
}
