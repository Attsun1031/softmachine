package config

import (
	"github.com/spf13/viper"
)

// Reference to https://polidog.jp/2016/08/09/golang_gorm1/

type DbConfig struct {
	User     string
	Password string
	Host     string
	Port     int
}

func LoadDbConfig() *DbConfig {
	dbConfig := &DbConfig{
		User:     viper.GetString("db.user"),
		Password: viper.GetString("db.password"),
		Host:     viper.GetString("db.host"),
		Port:     viper.GetInt("db.port"),
	}
	return dbConfig
}
