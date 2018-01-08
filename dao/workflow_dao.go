package dao

import (
	"github.com/Attsun1031/jobnetes/model"
	"github.com/jinzhu/gorm"
)

type WorkflowDao interface {
	FindById(uint, *gorm.DB) *model.Workflow
}

type WorkflowDaoImpl struct{}

func (dao *WorkflowDaoImpl) FindById(id uint, db *gorm.DB) *model.Workflow {
	wf := &model.Workflow{}
	db.Where("id = ?", id).First(wf)
	return wf
}
