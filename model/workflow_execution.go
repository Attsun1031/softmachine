package model

import "time"

type WorkflowStatusType int

const (
	WfScheduled WorkflowStatusType = iota
	WfRunning
	WfPending
	WfWaitingRetry
	WfCaceled
	WfSuccess
	WfFailed
)

var UncompletedWorkflowStatuses = []WorkflowStatusType{
	WfRunning,
	WfScheduled,
	WfPending,
	WfWaitingRetry,
}

type WorkflowExecution struct {
	ID         uint `gorm:"primary_key"`
	Workflow   *Workflow
	WorkflowID uint   `gorm:"not null"`
	Name       string `gorm:"not null"`
	StartedAt  *time.Time
	EndedAt    *time.Time
	Status     WorkflowStatusType `gorm:"not null"`
	Input      string             `gorm:"type:json;not null"`
	Output     string             `gorm:"type:json;not null"`
	RetryCount int                `gorm:"not null;default:0"`
	Errors     []ExecutionError   `gorm:"type:json;not null"`
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
}
