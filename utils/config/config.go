package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type _JobnetesConfig struct {
	DbConfig         *DbConfig
	LogConfig        *LogConfig
	KubernetesConfig *KubernetesConfig
}

var JobnetesConfig *_JobnetesConfig

func InitConfig() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/jobnetes")
	viper.AddConfigPath("$HOME/.jobnetes")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read config: %s \n", err))
	}

	JobnetesConfig = &_JobnetesConfig{
		DbConfig:         LoadDbConfig(),
		LogConfig:        LoadLogConfig(),
		KubernetesConfig: LoadKubernetesConfig(),
	}
}
