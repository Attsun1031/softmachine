package taskpoller

import (
	"errors"

	"fmt"

	"github.com/Attsun1031/jobnetes/dao"
	"github.com/Attsun1031/jobnetes/model"
	"k8s.io/client-go/kubernetes"
)

type Factory interface {
	GetTaskPoller(*model.TaskExecution, *model.JobDef) (TaskPoller, error)
}

type FactoryImpl struct {
	TaskExecutionDao dao.TaskExecutionDao
	KubeClient       kubernetes.Interface
}

func (factory *FactoryImpl) GetTaskPoller(te *model.TaskExecution, jobDef *model.JobDef) (TaskPoller, error) {
	currentJobDef := jobDef.GetCurrentTask(te)
	if currentJobDef == nil {
		return nil, errors.New(fmt.Sprintf("Current task not found in jobdef. taskName=%v", te.TaskName))
	}
	switch currentJobDef.GetJobType() {
	case model.JobTypeKube:
		return &KubeJobTaskPoller{
			Client:           factory.KubeClient,
			TaskExecutionDao: factory.TaskExecutionDao,
		}, nil
	case model.JobTypeParallel:
		return &ParallelJobTaskPoller{
			TaskExecutionDao: factory.TaskExecutionDao,
			TaskDef:          currentJobDef.(*model.ParallelTask),
		}, nil
	default:
		return nil, errors.New(fmt.Sprintf("Unknown task type %v", currentJobDef.GetJobType()))
	}
}
