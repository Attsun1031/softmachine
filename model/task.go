package model

type Task interface {
	GetName() string
	GetJobType() string
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
