package manager

import (
	"github.com/Attsun1031/jobnetes/dao/db"
	"github.com/Attsun1031/jobnetes/kubernetes"
	"github.com/Attsun1031/jobnetes/utils/log"
)

type _ManagerConfig struct {
	DbConfig   *db.DbConfig
	LogConfig  *log.LogConfig
	KubeConfig *kubernetes.KubeConfig
}

var ManagerConfig *_ManagerConfig

func load() *_ManagerConfig {
	ManagerConfig = &_ManagerConfig{
		DbConfig:   db.LoadDbConfig(),
		LogConfig:  log.LoadLogConfig(),
		KubeConfig: kubernetes.LoadKubeConfig(),
	}
	return ManagerConfig
}

func InitConfig() {
	ManagerConfig = load()
	log.SetupLogger(ManagerConfig.LogConfig)
}
