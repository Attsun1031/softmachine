package model

import "k8s.io/api/batch/v1"

type Task interface {
	GetName() string
	GetJobType() string
}

type KubeJobTask struct {
	Name         string
	NextTaskName string
	KubeJobSpec  v1.Job
}

func (task *KubeJobTask) GetName() string {
	return task.Name
}

func (task *KubeJobTask) GetJobType() string {
	return "kubejob"
}
