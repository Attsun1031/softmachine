package dao

import (
	"github.com/Attsun1031/jobnetes/model"
	"github.com/jinzhu/gorm"
)

type WorkflowExecutionDao interface {
	FindById(uint, *gorm.DB) *model.WorkflowExecution
	Update(*model.WorkflowExecution, *gorm.DB)
	FindUncompletedWorkflowExecs(*gorm.DB) []*model.WorkflowExecution
}

type WorkflowExecutionDaoImpl struct{}

func (workflowExecutionDaoImpl *WorkflowExecutionDaoImpl) FindById(id uint, db *gorm.DB) *model.WorkflowExecution {
	execution := &model.WorkflowExecution{}
	db.Where("id = ?", id).First(execution)
	return execution
}

func (workflowExecutionDaoImpl *WorkflowExecutionDaoImpl) Update(execution *model.WorkflowExecution, db *gorm.DB) {
	db.Save(execution)
}

func (workflowExecutionDaoImpl *WorkflowExecutionDaoImpl) FindUncompletedWorkflowExecs(db *gorm.DB) []*model.WorkflowExecution {
	var workflowExecutions []*model.WorkflowExecution
	db.
		Preload("Workflow").
		Where("status in (?)", model.UncompletedWorkflowStatuses).
		Find(&workflowExecutions)
	return workflowExecutions
}
