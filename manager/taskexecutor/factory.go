package taskexecutor

import (
	"errors"

	"fmt"
	"reflect"

	"github.com/Attsun1031/jobnetes/dao"
	"github.com/Attsun1031/jobnetes/model"
	"k8s.io/client-go/kubernetes"
)

type Factory interface {
	GetTaskExecutor(task model.Task) (TaskExecutor, error)
}

type FactoryImpl struct {
	TaskExecutionDao dao.TaskExecutionDao
	KubeClient       kubernetes.Interface
}

func (factory *FactoryImpl) GetTaskExecutor(task model.Task) (TaskExecutor, error) {
	switch task.(type) {
	case *model.KubeJobTask:
		return &KubeJobTaskExecutor{
			Task:             task.(*model.KubeJobTask),
			TaskExecutionDao: factory.TaskExecutionDao,
			KubeClient:       factory.KubeClient,
		}, nil
	default:
		return nil, errors.New(fmt.Sprintf("Unknown task type %s", reflect.TypeOf(task)))
	}
}
