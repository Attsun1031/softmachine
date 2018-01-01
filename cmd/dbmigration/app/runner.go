package app

import (
	"github.com/Attsun1031/jobnetes/dao/db"
	"github.com/Attsun1031/jobnetes/model"
	"github.com/Attsun1031/jobnetes/utils/config"
)

func Run() {
	config.InitConfig()
	d := db.Connect(config.JobnetesConfig.DbConfig)
	// TODO: foreign key
	// https://github.com/jinzhu/gorm/issues/450
	d.AutoMigrate(&model.User{}, &model.Workflow{}, &model.WorkflowExecution{}, &model.TaskExecution{})
}
