package model

import "time"

type TaskStatusType = int

const (
	TaskRunning TaskStatusType = iota
	TaskCaceled
	TaskSuccess
	TaskFailed
)

type TaskExecution struct {
	ID                  uint `gorm:"primary_key"`
	WorkflowExecution   *WorkflowExecution
	WorkflowExecutionID uint       `gorm:"not null"`
	Name                string     `gorm:"not null"`
	StartedAt           *time.Time `gorm:"not null"`
	EndedAt             *time.Time
	Status              TaskStatusType `gorm:"not null"`
	Input               string         `gorm:"type:json;not null"`
	Output              string         `gorm:"type:json;not null"`
	Attempt             int            `gorm:"not null;default:1"`
	ErrorMsg            string
	CreatedAt           *time.Time
	UpdatedAt           *time.Time
}
