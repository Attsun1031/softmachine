package dao

import (
	"github.com/Attsun1031/jobnetes/model"
	"github.com/jinzhu/gorm"
)

type WorkflowDao interface {
	FindById(uint, *gorm.DB) (*model.Workflow, error)
	FindByIds([]uint, *gorm.DB) ([]*model.Workflow, error)
}

type WorkflowDaoImpl struct{}

func (dao *WorkflowDaoImpl) FindById(id uint, db *gorm.DB) (*model.Workflow, error) {
	wf := &model.Workflow{}
	err := db.Where("id = ?", id).First(wf).Error
	return wf, err
}

func (dao *WorkflowDaoImpl) FindByIds(ids []uint, db *gorm.DB) ([]*model.Workflow, error) {
	var wfs []*model.Workflow
	err := db.Where("id in (?)", ids).Order("id").Find(&wfs).Error
	return wfs, err
}
