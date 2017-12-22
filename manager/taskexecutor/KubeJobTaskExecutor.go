package taskexecutor

import (
	"fmt"

	"github.com/Attsun1031/jobnetes/model"
	"github.com/Attsun1031/jobnetes/utils/log"
	"github.com/jinzhu/gorm"
)

type KubeJobTaskExecutor struct {
	Task *model.KubeJobTask
}

func (executor *KubeJobTaskExecutor) Execute(wfExecution *model.WorkflowExecution, db *gorm.DB) error {
	task := executor.Task
	log.Logger.Info(fmt.Sprintf("Requesting task. Name=%s Type=%s", task.Name, task.GetJobType()))

	// task_executionレコードを作成（トランザクションハンドリングは呼び出し側に任せる）

	// kubejobを起動
	return nil
}
