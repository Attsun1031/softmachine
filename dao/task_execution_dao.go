package dao

import (
	"github.com/Attsun1031/jobnetes/model"
	"github.com/jinzhu/gorm"
)

type TaskExecutionDao interface {
	FindById(uint, *gorm.DB) *model.TaskExecution
	FindCompletedByWorkflowId(uint, *gorm.DB) []*model.TaskExecution
	FindUncompletedByWorkflowId(uint, *gorm.DB) []*model.TaskExecution
	Update(*model.TaskExecution, *gorm.DB)
}

type TaskExecutionDaoImpl struct{}

func (dao *TaskExecutionDaoImpl) FindById(id uint, db *gorm.DB) *model.TaskExecution {
	execution := &model.TaskExecution{}
	db.Where("id = ?", id).First(execution)
	return execution
}

func (dao *TaskExecutionDaoImpl) FindCompletedByWorkflowId(wid uint, db *gorm.DB) []*model.TaskExecution {
	var executions []*model.TaskExecution
	db.
		Where("workflow_execution_id = ? AND status in (?)", wid, model.CompletedTaskStatuses).
		Find(&executions)
	return executions
}

func (dao *TaskExecutionDaoImpl) FindUncompletedByWorkflowId(wid uint, db *gorm.DB) []*model.TaskExecution {
	var executions []*model.TaskExecution
	db.
		Where("workflow_execution_id = ? AND status in (?)", wid, model.UncompletedTaskStatuses).
		Find(&executions)
	return executions
}

func (dao *TaskExecutionDaoImpl) Update(execution *model.TaskExecution, db *gorm.DB) {
	db.Save(execution)
}
