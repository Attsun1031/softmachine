package config

import "github.com/spf13/viper"

type JobApiConfig struct {
	Port     uint
	Username string
	Password string
}

func LoadJobApiConfig() *JobApiConfig {
	return &JobApiConfig{
		Port:     uint(viper.GetInt("jobapi.port")),
		Username: viper.GetString("jobapi.username"),
		Password: viper.GetString("jobapi.password"),
	}
}
