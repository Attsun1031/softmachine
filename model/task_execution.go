package model

import (
	"time"

	"github.com/Attsun1031/jobnetes/utils/config"
	"k8s.io/api/batch/v1"
	v1meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type TaskStatusType = int

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
	ID                  uint `gorm:"primary_key"`
	WorkflowExecution   *WorkflowExecution
	WorkflowExecutionID uint       `gorm:"not null"`
	ExecutionName       string     `gorm:"not null"`
	TaskName            string     `gorm:"not null"`
	StartedAt           *time.Time `gorm:"not null"`
	EndedAt             *time.Time
	Status              TaskStatusType `gorm:"not null"`
	Input               string         `gorm:"type:json;not null"`
	Output              string         `gorm:"type:json;not null"`
	ErrorReason         string
	ErrorMsg            string
	CreatedAt           *time.Time
	UpdatedAt           *time.Time
}

func (te *TaskExecution) Poll(client kubernetes.Interface) (bool, error) {
	result, err := client.
		BatchV1().
		Jobs(config.JobnetesConfig.KubernetesConfig.JobNamespace).
		Get(te.ExecutionName, v1meta.GetOptions{})

	if err != nil {
		return false, err
	}
	if len(result.Status.Conditions) == 0 {
		return false, nil
	}

	lastCondition := result.Status.Conditions[0]

	switch lastCondition.Type {
	case v1.JobComplete:
		te.Status = TaskSuccess
		te.EndedAt = &result.Status.CompletionTime.Time
	default:
		te.Status = TaskFailed
		te.EndedAt = &lastCondition.LastProbeTime.Time
		te.ErrorReason = lastCondition.Reason
		te.ErrorMsg = lastCondition.Message
	}
	return true, nil
}
