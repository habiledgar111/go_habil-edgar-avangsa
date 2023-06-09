package repository

import (
	"errors"
	"sec_23/Praktikum/restAPI_unit_testing/model"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (um *UserRepositoryMock) CreateUser(user *model.User) error {
	if user == nil {
		return errors.New("error")
	} else {
		return nil
	}
}
