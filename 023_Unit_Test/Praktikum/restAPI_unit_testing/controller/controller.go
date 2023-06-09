package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sec_23/Praktikum/restAPI_unit_testing/config"
	"sec_23/Praktikum/restAPI_unit_testing/controller/repository"
	"sec_23/Praktikum/restAPI_unit_testing/model"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/labstack/echo/v4"
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

type Controller struct {
}

func (m *Controller) GetUser(c echo.Context) error {
	user := c.Get("user").(model.UserMock)
	log.Println("user data : ", user)

	var users []model.UserMock

	config.DBMysql.Find(&users)
	return c.JSON(http.StatusOK, users)
}

func (m *Controller) CreateUser(c echo.Context) error {
	data := map[string]interface{}{
		"Message": "fail",
	}
	var user model.UserMock
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, data)
	}

	err = repository.GetUserRepository().CreateUser(&user)
	if err != nil {
		return err
	}

	data["massage"] = "success"

	return c.JSON(http.StatusOK, data)
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

// func GetUser(c echo.Context) error {
// 	var users model.User
// 	id := c.Param("id")
// 	if err := DB.First(&users, id).Error; err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "success get user",
// 		"user":    users,
// 	})
// }

func GetUser(c echo.Context) error {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"massage": "missing token",
		})
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"massage": "failed cast claims",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "get data",
		"user":    claims,
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
	temp := DB.Unscoped().Delete(user, id)
	if temp.RowsAffected < 1 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"massage": "cant find data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "success delete data",
	})
}

func UserLogin(ctx echo.Context) error {
	var user model.User
	json_map := make(map[string]interface{})
	err := json.NewDecoder(ctx.Request().Body).Decode(&json_map)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"Massage": "json cant empty",
		})
	}

	email := fmt.Sprintf("%v", json_map["email"])
	password := fmt.Sprintf("%v", json_map["password"])

	if err := DB.Where("email = @email", sql.Named("email", email)).First(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if user.Password != password {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"massage": "password salah",
		})
	}

	token, claims := CreateJWT(user)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"massage": "berhasil login",
		"token":   token,
		"users":   user,
		"claims":  claims,
	})
}

type jwtCustomClaims struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	jwt.RegisteredClaims
}

func CreateJWT(user model.User) (interface{}, interface{}) {
	email := user.Email
	pass := user.Password
	claims := &jwtCustomClaims{
		email,
		pass,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	temp := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := temp.SignedString([]byte("secret"))

	if err != nil {
		return err.Error(), nil
	}

	return token, claims
}
