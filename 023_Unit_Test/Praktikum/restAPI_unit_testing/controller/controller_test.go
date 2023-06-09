package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"sec_23/Praktikum/restAPI_unit_testing/controller/repository"
	"sec_23/Praktikum/restAPI_unit_testing/model"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

func TestCreateUser(t *testing.T) {
	userRepo := &repository.UserRepositoryMock{Mock: mock.Mock{}}

	repository.SetUserRepository(userRepo)

	type args struct {
		c echo.Context
	}

	test := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "success",
			wantErr: false,
		},
	}
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			data := model.UserMock{
				Email:    "emailmock111@gmail.com",
				Password: "passwordmock",
				Name:     "Namemock",
				Age:      99,
			}
			userRepo.Mock.On("CreateUser", &data).Return(errors.New("error"))

			e := echo.New()

			bData, _ := json.Marshal(data)
			req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(bData))
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			controller := Controller{}
			controller.CreateUser(c)

		})
	}
}
