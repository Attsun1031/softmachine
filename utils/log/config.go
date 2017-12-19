package log

import (
	"github.com/Attsun1031/jobnetes/utils/consts"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

type LogConfig struct {
	LogLevel logrus.Level
}

const defaultLogLevel = logrus.InfoLevel

func LoadLogConfig() *LogConfig {
	return &LogConfig{
		LogLevel: getLevel(),
	}
}

func getLevel() logrus.Level {
	levelStr := strings.ToLower(os.Getenv(consts.LogLevel))
	for _, l := range logrus.AllLevels {
		if levelStr == l.String() {
			return l
		}
	}
	return defaultLogLevel
}
