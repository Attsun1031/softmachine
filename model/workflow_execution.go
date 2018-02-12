package model

import "time"

type WorkflowStatusType int

const (
	WfScheduled WorkflowStatusType = iota
	WfRunning
	WfSuccess
	WfFailed
)

var UncompletedWorkflowStatuses = []WorkflowStatusType{
	WfRunning,
	WfScheduled,
}

type WorkflowExecution struct {
	ID             uint               `gorm:"primary_key" json:"id"`
	Workflow       *Workflow          `json:"workflow"`
	WorkflowID     uint               `gorm:"not null" json:"-"`
	TaskExecutions []TaskExecution    `json:"taskExecutions"`
	Name           string             `gorm:"not null" json:"name"`
	StartedAt      *time.Time         `json:"startedAt"`
	EndedAt        *time.Time         `json:"endedAt"`
	Status         WorkflowStatusType `gorm:"not null" json:"status"`
	Input          string             `gorm:"type:json;not null" json:"input"`
	Output         string             `gorm:"type:json;not null" json:"output"`
	Definition     string             `gorm:"type:json;not null" json:"definition"`
	Attempt        int                `gorm:"not null;default:1" json:"attempt"`
	ErrorMsg       string             `json:"errorMsg"`
	CreatedAt      *time.Time         `json:"-"`
	UpdatedAt      *time.Time         `json:"-"`

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
