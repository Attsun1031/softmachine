package dao

import (
	"github.com/Attsun1031/jobnetes/model"
	"github.com/jinzhu/gorm"
)

type WorkflowExecutionDao interface {
	FindById(uint, bool, bool, *gorm.DB) (*model.WorkflowExecution, error)
	Find(limit int, offset int, order string, db *gorm.DB) ([]*model.WorkflowExecution, error)
	Create(*model.WorkflowExecution, *gorm.DB) error
	Update(*model.WorkflowExecution, *gorm.DB) error
	FindUncompletedWorkflowExecs(*gorm.DB) ([]*model.WorkflowExecution, error)
}

type WorkflowExecutionDaoImpl struct{}

func (workflowExecutionDaoImpl *WorkflowExecutionDaoImpl) FindById(id uint, withWorkflow bool, withTasks bool, db *gorm.DB) (*model.WorkflowExecution, error) {
	execution := &model.WorkflowExecution{}
	var d *gorm.DB
	if withTasks {
		d = db.Preload("TaskExecutions")
	} else {
		d = db
	}
	if withWorkflow {
		d = d.Preload("Workflow")
	} else {
		d = d
	}
	err := d.Where("id = ?", id).First(execution).Error
	return execution, err
}

func (workflowExecutionDaoImpl *WorkflowExecutionDaoImpl) Find(limit int, offset int, order string, db *gorm.DB) ([]*model.WorkflowExecution, error) {
	var workflowExecutions []*model.WorkflowExecution
	err := db.
		Order(order).
		Limit(limit).
		Offset(offset).
		Find(&workflowExecutions).
		Error
	return workflowExecutions, err
}

func (workflowExecutionDaoImpl *WorkflowExecutionDaoImpl) Create(execution *model.WorkflowExecution, db *gorm.DB) error {
	if execution.Input == "" {
		execution.Input = "{}"
	}
	if execution.Output == "" {
		execution.Output = "{}"
	}
	return db.Create(execution).Error
}

func (workflowExecutionDaoImpl *WorkflowExecutionDaoImpl) Update(execution *model.WorkflowExecution, db *gorm.DB) error {
	return db.Save(execution).Error
}

func (workflowExecutionDaoImpl *WorkflowExecutionDaoImpl) FindUncompletedWorkflowExecs(db *gorm.DB) ([]*model.WorkflowExecution, error) {
	var workflowExecutions []*model.WorkflowExecution
	err := db.
		Preload("Workflow").
		Where("status in (?)", model.UncompletedWorkflowStatuses).
		Find(&workflowExecutions).
		Error
	return workflowExecutions, err
}
