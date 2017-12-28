package db

import (
	"fmt"

	"github.com/Attsun1031/jobnetes/utils/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Reference to https://polidog.jp/2016/08/09/golang_gorm1/

func Connect(config *config.DbConfig) *gorm.DB {
	var url string
	if config.Password == "" {
		url = fmt.Sprintf("%s@tcp(%s:%d)/jobnetes?charset=utf8&parseTime=True&loc=Local", config.User, config.Host, config.Port)
	}
	db, err := gorm.Open("mysql", url)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect database. message = %s", err))
	}
	return db
}
