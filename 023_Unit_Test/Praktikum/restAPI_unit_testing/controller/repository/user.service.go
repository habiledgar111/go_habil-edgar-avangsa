package repository

import (
	"sec_23/Praktikum/restAPI_unit_testing/config"
	"sec_23/Praktikum/restAPI_unit_testing/model"
)

type IUserService interface {
	CreateUser(*model.User) error
}

type UserRepository struct {
	Func IUserService
}

var userRepository IUserService

func init() {
	ur := &UserRepository{}
	ur.Func = ur

	userRepository = ur
}

func GetUserRepository() IUserService {
	return userRepository
}

func SetUserRepository(ur IUserService) {
	userRepository = ur
}

func (u *UserRepository) CreateUser(user *model.User) error {
	err := config.DBMysql.Save(&user)
	if err != nil {
		return err.Error
	}

	return nil
}
