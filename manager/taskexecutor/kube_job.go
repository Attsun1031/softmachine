package taskexecutor

import (
	"fmt"

	"time"

	"github.com/Attsun1031/jobnetes/dao"
	"github.com/Attsun1031/jobnetes/model"
	"github.com/Attsun1031/jobnetes/utils/log"
	"github.com/jinzhu/gorm"
	"k8s.io/api/batch/v1"
	"k8s.io/client-go/kubernetes"
)

type KubeJobTaskExecutor struct {
	Task             *model.KubeJobTask
	TaskExecutionDao dao.TaskExecutionDao
	KubeClient       kubernetes.Interface
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
		Output:            "{}",
	}
	executor.TaskExecutionDao.Update(te, db)

	// start kubejob
	// k8s.io/client-go/kubernetes/typed/batch/v1/job.go#Createのパクリ
	result := &v1.Job{}
	return executor.KubeClient.BatchV1().RESTClient().Post().
		Namespace("default").
		Resource("jobs").
		Body([]byte(task.KubeJobSpec)).
		Do().
		Into(result)
}
