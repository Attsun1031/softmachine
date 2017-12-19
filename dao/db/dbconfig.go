package db

import (
	"os"

	"github.com/Attsun1031/jobnetes/utils/consts"
)

// Reference to https://polidog.jp/2016/08/09/golang_gorm1/

type DbConfig struct {
	Url string
}

const defaultDbUrl = "root@tcp(127.0.0.1:3333)/jobnetes?charset=utf8&parseTime=True&loc=Local"

func LoadDbConfig() *DbConfig {
	dbConfig := &DbConfig{
		defaultDbUrl,
	}
	url := os.Getenv(consts.DbUrl)
	if url != "" {
		dbConfig.Url = url
	}
	return dbConfig
}
