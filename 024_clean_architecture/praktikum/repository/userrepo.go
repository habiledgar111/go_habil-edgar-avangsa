package repository

import (
	"errors"
	"sec024/praktikum/config"
	"sec024/praktikum/entity"
)

var (
	user *entity.User
)

func GetAllUser() (entity.User, error) {
	err := config.DB.Find(user).Error

	if err != nil {
		return *user, errors.New("error massage")
	}

	return *user, nil
}

func Createuser(usernew entity.User) error {
	err := config.DB.Create(usernew).Error
	if err != nil {
		return err
	}
	return nil
}
