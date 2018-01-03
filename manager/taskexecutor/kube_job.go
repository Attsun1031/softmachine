package taskexecutor

import (
	"fmt"

	"time"

	"github.com/Attsun1031/jobnetes/dao"
	"github.com/Attsun1031/jobnetes/model"
	"github.com/Attsun1031/jobnetes/utils/config"
	"github.com/Attsun1031/jobnetes/utils/log"
	"github.com/jinzhu/gorm"
	"k8s.io/client-go/kubernetes"
)

type KubeJobTaskExecutor struct {
	Task             *model.KubeJobTask
	TaskExecutionDao dao.TaskExecutionDao
	KubeClient       kubernetes.Interface
}

func (executor *KubeJobTaskExecutor) Execute(we *model.WorkflowExecution, db *gorm.DB, input string) error {
	task := executor.Task
	log.Logger.Infof("Requesting task. ExecutionName=%s Type=%s", task.Name, task.GetJobType())

	// create execution record
	startedAt := time.Now()
	te := &model.TaskExecution{
		WorkflowExecution: we,
		TaskName:          task.Name,
		StartedAt:         &startedAt,
		Status:            model.TaskRunning,
		Input:             input,
		Output:            "{}",
	}
	executor.TaskExecutionDao.Update(te, db)
	name := fmt.Sprintf(
		"%s-%d-%d-%s",
		task.Name,
		we.ID,
		te.ID,
		time.Now().Format("2006-01-02-15-04-05-99"))
	te.ExecutionName = name
	executor.TaskExecutionDao.Update(te, db)

	// start kubernetes job
	spec := task.KubeJobSpec
	spec.Name = name
	_, err := executor.KubeClient.
		BatchV1().
		Jobs(config.JobnetesConfig.KubernetesConfig.JobNamespace).
		Create(&spec)
	return err
}
