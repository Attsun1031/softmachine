package dao

import (
	"github.com/Attsun1031/jobnetes/model"
	"github.com/jinzhu/gorm"
)

type WorkflowDao interface {
	FindById(uint, *gorm.DB) (*model.Workflow, error)
}

type WorkflowDaoImpl struct{}

func (dao *WorkflowDaoImpl) FindById(id uint, db *gorm.DB) (*model.Workflow, error) {
	wf := &model.Workflow{}
	err := db.Where("id = ?", id).First(wf).Error
	return wf, err
}
