package dao

import (
	"github.com/Attsun1031/jobnetes/model"
	"github.com/jinzhu/gorm"
)

type TaskExecutionDao interface {
	Update(*model.TaskExecution, *gorm.DB)
}

type TaskExecutionDaoImpl struct{}

func (dao *TaskExecutionDaoImpl) Update(execution *model.TaskExecution, db *gorm.DB) {
	db.Save(execution)
}
