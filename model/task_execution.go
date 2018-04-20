package model

import (
	"time"
	"math"
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
	ID                    uint               `gorm:"primary_key" json:"id"`
	WorkflowExecution     *WorkflowExecution `json:"-"`
	WorkflowExecutionID   uint               `gorm:"not null" json:"-"`
	ParentTaskExecution   *TaskExecution     `json:"-"`
	ParentTaskExecutionID uint               `json:"parentId"`
	NextTaskExecution     *TaskExecution     `json:"-"`
	NextTaskExecutionID   uint               `json:"nextId"`
	PrevTaskExecution     *TaskExecution     `json:"-"`
	PrevTaskExecutionID   uint               `json:"prevId"`
	ExecutionName         string             `gorm:"not null" json:"executionName"`
	TaskName              string             `gorm:"not null" json:"taskName"`
	TaskType              string             `gorm:"not null" json:"taskType"`
	StartedAt             *time.Time         `gorm:"not null" json:"startedAt"`
	EndedAt               *time.Time         `json:"endedAt"`
	ElapsedSec            uint               `json:"elapsedSec"`
	Status                TaskStatusType     `gorm:"not null" json:"status"`
	Input                 string             `gorm:"type:json;not null" json:"input"`
	Output                string             `gorm:"type:json;not null" json:"output"`
	ErrorReason           string             `json:"errorReason"`
	ErrorMsg              string             `json:"errorMsg"`
	CreatedAt             *time.Time         `json:"-"`
	UpdatedAt             *time.Time         `json:"-"`
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

func (te *TaskExecution) MarkSuccess(endedAt *time.Time) {
	te.Status = TaskSuccess
	te.setEndTime(endedAt)
}

func (te *TaskExecution) MarkFailed(endedAt *time.Time, errReason string, errMsg string) {
	te.Status = TaskFailed
	te.ErrorReason = errReason
	te.ErrorMsg = errMsg
	te.setEndTime(endedAt)
}

func (te *TaskExecution) setEndTime(endedAt *time.Time) {
	te.EndedAt = endedAt
	te.ElapsedSec = uint(math.Ceil(endedAt.Sub(*te.StartedAt).Seconds()))
}
