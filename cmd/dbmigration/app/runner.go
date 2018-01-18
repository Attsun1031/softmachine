package app

import (
	"github.com/Attsun1031/jobnetes/dao/db"
	"github.com/Attsun1031/jobnetes/model"
	"github.com/Attsun1031/jobnetes/utils/config"
	"github.com/Attsun1031/jobnetes/utils/log"
)

func Run() {
	config.InitConfig()
	log.SetupLogger(config.JobnetesConfig.LogConfig)
	d := db.Connect(config.JobnetesConfig.DbConfig)
	// TODO: foreign key
	// https://github.com/jinzhu/gorm/issues/450
	d.Exec("CREATE DATABASE IF NOT EXISTS jobnetes;")
	d.AutoMigrate(&model.User{}, &model.Workflow{}, &model.WorkflowExecution{}, &model.TaskExecution{})
}
