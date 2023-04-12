package config

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBMysql *gorm.DB
)

func InitMockDB() (*sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, err
	}

	dialector := mysql.New(mysql.Config{
		DSN:                       "sqlmock_db",
		DriverName:                "mysql",
		Conn:                      db,
		SkipInitializeWithVersion: true,
	})

	DBMysql, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &mock, nil
}
