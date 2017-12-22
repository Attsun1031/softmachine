package taskexecutor

import (
	"fmt"

	"time"

	"github.com/Attsun1031/jobnetes/dao"
	"github.com/Attsun1031/jobnetes/model"
	"github.com/Attsun1031/jobnetes/utils/log"
	"github.com/jinzhu/gorm"
	"k8s.io/client-go/rest"
)

type KubeJobTaskExecutor struct {
	Task             *model.KubeJobTask
	TaskExecutionDao dao.TaskExecutionDao
}

func (executor *KubeJobTaskExecutor) Execute(we *model.WorkflowExecution, db *gorm.DB) error {
	task := executor.Task
	log.Logger.Info(fmt.Sprintf("Requesting task. Name=%s Type=%s", task.Name, task.GetJobType()))

	// task_executionレコードを作成（トランザクションハンドリングは呼び出し側に任せる）
	startedAt := time.Now()
	te := &model.TaskExecution{
		WorkflowExecution: we,
		Name:              task.Name,
		StartedAt:         &startedAt,
		Status:            model.TaskRunning,
		Input:             we.Input,
	}
	executor.TaskExecutionDao.Update(te, db)

	// kubejobを起動
	config, err := rest.InClusterConfig()

	return nil
}
