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

func (executor *KubeJobTaskExecutor) Execute(we *model.WorkflowExecution, db *gorm.DB, input string, parentID uint) error {
	task := executor.Task
	log.Logger.Infof("Requesting k8s task. ExecutionName=%s Type=%s ParentID=%v", task.Name, task.GetJobType(), parentID)

	// create execution record
	startedAt := time.Now()
	te := &model.TaskExecution{
		WorkflowExecution:     we,
		ParentTaskExecutionID: parentID,
		TaskName:              task.Name,
		TaskType:              task.GetJobType(),
		StartedAt:             &startedAt,
		Status:                model.TaskRunning,
		Input:                 input,
		Output:                "{}",
	}
	err := executor.TaskExecutionDao.Update(te, db)
	if err != nil {
		return err
	}

	name := fmt.Sprintf(
		"%s-%d-%d-%s",
		task.Name,
		we.ID,
		te.ID,
		time.Now().Format("2006-01-02-15-04-05-99"))
	te.ExecutionName = name
	err = executor.TaskExecutionDao.Update(te, db)
	if err != nil {
		return err
	}

	// start kubernetes job
	spec := task.KubeJobSpec
	spec.Name = name
	_, err = executor.KubeClient.
		BatchV1().
		Jobs(config.JobnetesConfig.KubernetesConfig.JobNamespace).
		Create(&spec)
	return err
}
