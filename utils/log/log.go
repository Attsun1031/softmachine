package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Logger *logrus.Logger
var ErrLogger = &logrus.Logger{
	Out:       os.Stderr,
	Formatter: new(logrus.JSONFormatter),
	Level:     logrus.ErrorLevel,
}

func SetupLogger(config *LogConfig) {
	levels := []logrus.Level{logrus.ErrorLevel, logrus.PanicLevel, logrus.FatalLevel}
	hook := NewMultiLogHook(levels, ErrLogger)
	Logger = &logrus.Logger{
		Out:       os.Stdout,
		Formatter: new(logrus.JSONFormatter),
		Level:     config.LogLevel,
		Hooks:     make(logrus.LevelHooks),
	}
	Logger.AddHook(hook)
}
