package db

import (
	"fmt"

	"github.com/Attsun1031/jobnetes/utils/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
)

// Reference to https://polidog.jp/2016/08/09/golang_gorm1/

func Connect(config *config.DbConfig, logger *logrus.Logger) *gorm.DB {
	var userPass string
	if config.Password != "" {
		userPass = fmt.Sprintf("%s:%s", config.User, config.Password)
	} else {
		userPass = config.User
	}
	url := fmt.Sprintf("%s@tcp(%s:%d)/jobnetes?charset=utf8&parseTime=True&loc=Local", userPass, config.Host, config.Port)
	db, err := gorm.Open("mysql", url)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect database. message = %s", err))
	}
	if logger != nil {
		db.SetLogger(logger)
	}
	return db
}
