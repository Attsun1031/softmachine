package config

import "github.com/spf13/viper"

type WebAdminConfig struct {
	Port     uint
	Username string
	Password string
}

func LoadWebAdminConfig() *WebAdminConfig {
	return &WebAdminConfig{
		Port:     uint(viper.GetInt("webadmin.port")),
		Username: viper.GetString("webadmin.username"),
		Password: viper.GetString("webadmin.password"),
	}
}
