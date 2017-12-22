package model

import "time"

type WorkflowStatusType int

const (
	WfScheduled WorkflowStatusType = iota
	WfRunning
	WfPending
	WfCaceled
	WfSuccess
	WfFailed
)

var UncompletedWorkflowStatuses = []WorkflowStatusType{
	WfRunning,
	WfScheduled,
	WfPending,
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
	Definition string             `gorm:"type:json;not null"`
	Attempt    int                `gorm:"not null;default:1"`
	ErrorMsg   string
	CreatedAt  *time.Time
	UpdatedAt  *time.Time

	jobDef *JobDef `gorm:"-"`
}

type ExecutionError struct {
	cause   string
	message string
}

func (execution *WorkflowExecution) GetJobDef() *JobDef {
	if execution.jobDef != nil {
		return execution.jobDef
	}

	execution.jobDef = GetJobDefFromString(execution.Definition)
	return execution.jobDef
}
