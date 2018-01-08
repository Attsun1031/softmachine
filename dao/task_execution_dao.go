package dao

import (
	"github.com/Attsun1031/jobnetes/model"
	"github.com/jinzhu/gorm"
)

type TaskExecutionDao interface {
	FindById(uint, *gorm.DB) (*model.TaskExecution, error)
	FindCompletedByWorkflowId(uint, *gorm.DB) ([]*model.TaskExecution, error)
	FindUncompletedByWorkflowId(uint, *gorm.DB) ([]*model.TaskExecution, error)
	Update(*model.TaskExecution, *gorm.DB) error
}

type TaskExecutionDaoImpl struct{}

func (dao *TaskExecutionDaoImpl) FindById(id uint, db *gorm.DB) (*model.TaskExecution, error) {
	execution := &model.TaskExecution{}
	err := db.Where("id = ?", id).First(execution).Error
	return execution, err
}

func (dao *TaskExecutionDaoImpl) FindCompletedByWorkflowId(wid uint, db *gorm.DB) ([]*model.TaskExecution, error) {
	var executions []*model.TaskExecution
	err := db.
		Where("workflow_execution_id = ? AND status in (?)", wid, model.CompletedTaskStatuses).
		Find(&executions).
		Error
	return executions, err
}

func (dao *TaskExecutionDaoImpl) FindUncompletedByWorkflowId(wid uint, db *gorm.DB) ([]*model.TaskExecution, error) {
	var executions []*model.TaskExecution
	err := db.
		Where("workflow_execution_id = ? AND status in (?)", wid, model.UncompletedTaskStatuses).
		Find(&executions).
		Error
	return executions, err
}

func (dao *TaskExecutionDaoImpl) Update(execution *model.TaskExecution, db *gorm.DB) error {
	return db.Save(execution).Error
}
