package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

type MultiLogHook struct {
	ValidLevels []logrus.Level
	Logger      *logrus.Logger
}

func NewMultiLogHook(levels []logrus.Level, logger *logrus.Logger) *MultiLogHook {
	return &MultiLogHook{
		ValidLevels: levels,
		Logger:      logger,
	}
}

func (hook *MultiLogHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read entry, %v", err)
		return err
	}

	logger := hook.Logger
	switch entry.Level {
	case logrus.PanicLevel:
		logger.Panic(line)
	case logrus.FatalLevel:
		logger.Fatal(line)
	case logrus.ErrorLevel:
		logger.Error(line)
	case logrus.WarnLevel:
		logger.Error(line)
	case logrus.InfoLevel:
		logger.Info(line)
	case logrus.DebugLevel:
		logger.Debug(line)
	}
	return nil
}

func (hook *MultiLogHook) Levels() []logrus.Level {
	return hook.ValidLevels
}
