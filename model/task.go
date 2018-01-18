package model

import (
	"k8s.io/api/batch/v1"
)

type Task interface {
	GetName() string
	GetNextTaskName() string
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

func (task *KubeJobTask) GetNextTaskName() string {
	return task.NextTaskName
}

func (task *KubeJobTask) GetJobType() string {
	return JobTypeKube
}

type ParallelTask struct {
	Name         string
	NextTaskName string
	TaskSets     [][]Task
}

func (task *ParallelTask) GetName() string {
	return task.Name
}

func (task *ParallelTask) GetNextTaskName() string {
	return task.NextTaskName
}

func (task *ParallelTask) GetTaskSets() [][]Task {
	return task.TaskSets
}

func (task *ParallelTask) GetJobType() string {
	return JobTypeParallel
}
