package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"sec_22/022_Middleware/Pratikum/model"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	user *model.User
	book *model.Book
)

func init() {
	InitDB()
	InitMigration()
}

func InitMigration() {
	DB.AutoMigrate(user, book)
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
	email := fmt.Sprintf("%v", json_map["email"])
	password := fmt.Sprintf("%v", json_map["password"])
	users = model.User{
		Email:    email,
		Password: password,
		Name:     name,
		Age:      age,
	}
	result := DB.Create(&users)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":     "success create user",
		"rowaffected": result.RowsAffected,
		"querybool":   result.QueryFields,
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

func GetAllBooks(c echo.Context) error {
	var books []model.Book
	if err := DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   books,
	})
}

func GetBook(c echo.Context) error {
	var book model.Book
	id := c.Param("id")
	if err := DB.First(&book, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user",
		"book":    book,
	})
}

func CreateBook(c echo.Context) error {
	var books model.Book
	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Massage": "json cant empty",
		})
	}
	judul := fmt.Sprintf("%v", json_map["judul"])
	penulis := fmt.Sprintf("%v", json_map["penulis"])
	penerbit := fmt.Sprintf("%v", json_map["penerbit"])

	books = model.Book{
		Judul:    judul,
		Penulis:  penulis,
		Penerbit: penerbit,
	}
	result := DB.Create(&books)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":     "success create user",
		"rowaffected": result.RowsAffected,
	})
}

func UpdateBook(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	json_map := make(map[string]interface{})
	var dbbook model.Book

	if err := DB.First(&dbbook, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Massage": "json cant empty",
		})
	}

	if json_map["judul"] != nil {
		dbbook.Judul = fmt.Sprintf("%v", json_map["judul"])
	}
	if json_map["penulis"] != nil {
		dbbook.Penulis = fmt.Sprintf("%v", json_map["penulis"])
	}
	if json_map["penerbit"] != nil {
		dbbook.Penerbit = fmt.Sprintf("%v", json_map["penerbit"])
	}

	DB.Save(&dbbook)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "success update data",
		"user":    dbbook,
	})
}

func DeleteBook(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	temp := DB.Unscoped().Delete(book, id)
	if temp.RowsAffected < 1 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"massage": "cant find data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "success delete data",
	})
}

func UserLogin(c echo.Context) error {
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
	if err := DB.Where("email = @email", sql.Named("email", email)).First(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if users.Password != password {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"massage": "password salah",
		})
	}
	token := CreateJWT(users)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
		"user":  users,
	})
}

type jwtCustomClaims struct {
	email    string
	password string
	jwt.RegisteredClaims
}

func CreateJWT(user model.User) interface{} {
	claims := &jwtCustomClaims{
		user.Email,
		user.Password,
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
