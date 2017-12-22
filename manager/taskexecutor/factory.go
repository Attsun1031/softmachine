package taskexecutor

import (
	"errors"

	"fmt"
	"reflect"

	"github.com/Attsun1031/jobnetes/di"
	"github.com/Attsun1031/jobnetes/model"
)

type Factory interface {
	GetTaskExecutor(task model.Task) (TaskExecutor, error)
}

type FactoryImpl struct{}

func (factory *FactoryImpl) GetTaskExecutor(task model.Task) (TaskExecutor, error) {
	switch task.(type) {
	case *model.KubeJobTask:
		return &KubeJobTaskExecutor{
			Task:             task.(*model.KubeJobTask),
			TaskExecutionDao: di.InjectTaskExecutionDao(),
		}, nil
	default:
		return nil, errors.New(fmt.Sprintf("Unknown task type %s", reflect.TypeOf(task)))
	}
}
