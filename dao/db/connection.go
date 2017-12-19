package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Reference to https://polidog.jp/2016/08/09/golang_gorm1/

func Connect(config *DbConfig) *gorm.DB {
	db, err := gorm.Open("mysql", config.Url)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect database. message = %s", err))
	}
	return db
}
