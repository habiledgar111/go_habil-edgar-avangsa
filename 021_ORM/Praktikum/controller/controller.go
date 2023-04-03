package controller

import (
	"021_ORM/021_ORM/Praktikum/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	user *model.User
)

func init() {
	InitDB()
	InitMigration()
}

func InitMigration() {
	DB.AutoMigrate(user)
}

func InitDB() {
	dsn := "root:Mbahbambang123@tcp(localhost:3306)/sec21orm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
}

func SayHello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello word")
}

func GetAllUser(c echo.Context) error {
	var users []model.User
	if err := DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

func GetUser(c echo.Context) error {
	var users model.User
	id := c.Param("id")
	if err := DB.First(&users, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user",
		"user":    users,
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
	name := fmt.Sprintf("%v", json_map["name"])
	age, _ := strconv.Atoi(fmt.Sprintf("%v", json_map["age"]))
	users = model.User{
		Name: name,
		Age:  age,
	}
	result := DB.Create(&users)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":     "success create user",
		"rowaffected": result.RowsAffected,
	})
}

func UpdateUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	json_map := make(map[string]interface{})
	var dbuser model.User

	if err := DB.First(&dbuser, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Massage": "json cant empty",
		})
	}

	if json_map["name"] != nil {
		dbuser.Name = fmt.Sprintf("%v", json_map["name"])
	}
	if json_map["age"] != nil {
		dbuser.Age, _ = strconv.Atoi(fmt.Sprintf("%v", json_map["age"]))
	}

	DB.Save(&dbuser)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "success update data",
		"user":    dbuser,
	})
}

func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	DB.Delete(user, id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "success delete data",
	})
}
