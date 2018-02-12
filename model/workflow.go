package model

import (
	"time"
)

type Workflow struct {
	ID         uint       `gorm:"primary_key" json:"id"`
	Name       string     `gorm:"not null" json:"name"`
	Definition string     `gorm:"type:json;not null" json:"definition"`
	User       *User      `gorm:"not null" json:"user"`
	UserID     uint       `gorm:"not null" json:"-"`
	CreatedAt  *time.Time `json:"-"`
	UpdatedAt  *time.Time `json:"-"`

	jobDef *JobDef `gorm:"-"`
}

func (workflow *Workflow) GetJobDef() *JobDef {
	if workflow.jobDef != nil {
		return workflow.jobDef
	}

	workflow.jobDef = GetJobDefFromString(workflow.Definition)
	return workflow.jobDef
}
