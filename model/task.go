package model

import (
	"fmt"

	"github.com/Attsun1031/jobnetes/utils/log"
)

type Task interface {
	GetName() string
	GetJobType() string
	RequestTask(*WorkflowExecution) error
}

type KubeJobTask struct {
	Name         string
	NextTaskName string
	KubeJobSpec  string // TODO: Use Job in https://github.com/kubernetes/api/blob/master/batch/v1/types.go
}

func (task *KubeJobTask) GetName() string {
	return task.Name
}

func (task *KubeJobTask) GetJobType() string {
	return "kubejob"
}

func (task *KubeJobTask) RequestTask(execution *WorkflowExecution) error {
	log.Logger.Info(fmt.Sprintf("Requesting task. Name=%s Type=%s", task.Name, task.GetJobType()))
	return nil
}
