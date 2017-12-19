package app

import (
	"github.com/Attsun1031/jobnetes/model"
	"github.com/Attsun1031/jobnetes/model/db"
)

func Run() {
	dbConfig := db.LoadDbConfig()
	db.Connect(dbConfig)
	// TODO: foreign key
	// https://github.com/jinzhu/gorm/issues/450
	db.Db.AutoMigrate(&model.User{}, &model.Workflow{}, &model.WorkflowExecution{})
}
