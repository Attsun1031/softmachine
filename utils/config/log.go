package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type LogConfig struct {
	LogLevel logrus.Level
}

func LoadLogConfig() *LogConfig {
	return &LogConfig{
		LogLevel: getLevel(),
	}
}

func getLevel() logrus.Level {
	levelStr := viper.GetString("log.level")
	for _, l := range logrus.AllLevels {
		if levelStr == l.String() {
			return l
		}
	}
	panic("Log level not set.")
}
