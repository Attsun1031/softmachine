package model

import (
	"time"
)

type Workflow struct {
	ID         uint   `gorm:"primary_key"`
	Name       string `gorm:"not null"`
	Definition string `gorm:"type:json;not null"`
	User       *User  `gorm:"not null"`
	UserID     uint   `gorm:"not null"`
	CreatedAt  *time.Time
	UpdatedAt  *time.Time

	jobDef *JobDef `gorm:"-"`
}

func (workflow *Workflow) GetJobDef() *JobDef {
	if workflow.jobDef != nil {
		return workflow.jobDef
	}

	workflow.jobDef = GetJobDefFromString(workflow.Definition)
	return workflow.jobDef
}
