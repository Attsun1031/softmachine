package model

import (
	"time"
)

type TaskStatusType int

const (
	TaskRunning TaskStatusType = iota
	TaskSuccess
	TaskFailed
)

var UncompletedTaskStatuses = []TaskStatusType{
	TaskRunning,
}

var CompletedTaskStatuses = []TaskStatusType{
	TaskSuccess,
	TaskFailed,
}

type TaskExecution struct {
	ID                    uint `gorm:"primary_key"`
	WorkflowExecution     *WorkflowExecution
	WorkflowExecutionID   uint `gorm:"not null"`
	ParentTaskExecution   *TaskExecution
	ParentTaskExecutionID uint
	ExecutionName         string     `gorm:"not null"`
	TaskName              string     `gorm:"not null"`
	StartedAt             *time.Time `gorm:"not null"`
	EndedAt               *time.Time
	Status                TaskStatusType `gorm:"not null"`
	Input                 string         `gorm:"type:json;not null"`
	Output                string         `gorm:"type:json;not null"`
	ErrorReason           string
	ErrorMsg              string
	CreatedAt             *time.Time
	UpdatedAt             *time.Time
}

func (te *TaskExecution) IsCompleted() bool {
	switch te.Status {
	case TaskSuccess:
		return true
	case TaskFailed:
		return true
	default:
		return false
	}
}

func (te *TaskExecution) IsSucceeded() bool {
	return te.Status == TaskSuccess
}

func (te *TaskExecution) IsFailed() bool {
	return te.Status == TaskFailed
}
