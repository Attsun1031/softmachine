package app

import (
	"github.com/Attsun1031/jobnetes/dao/db"
	"github.com/Attsun1031/jobnetes/model"
)

func Run() {
	dbConfig := db.LoadDbConfig()
	db := db.Connect(dbConfig)
	// TODO: foreign key
	// https://github.com/jinzhu/gorm/issues/450
	db.AutoMigrate(&model.User{}, &model.Workflow{}, &model.WorkflowExecution{}, &model.TaskExecution{})
}
