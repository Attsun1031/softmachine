package workflowstate

import (
	"time"

	"fmt"

	"errors"

	"github.com/Attsun1031/jobnetes/dao"
	"github.com/Attsun1031/jobnetes/manager/taskexecutor"
	"github.com/Attsun1031/jobnetes/model"
	"github.com/Attsun1031/jobnetes/utils/log"
	"github.com/jinzhu/gorm"
	"k8s.io/client-go/kubernetes"
)

type RunningStateProcessor struct {
	WorkflowExecutionDao dao.WorkflowExecutionDao
	TaskExecutionDao     dao.TaskExecutionDao
	TaskExecutorFactory  taskexecutor.Factory
	KubeClient           kubernetes.Interface
}

// TODO: retry failed workflow
func (processor *RunningStateProcessor) ToNextState(execution *model.WorkflowExecution, db *gorm.DB) (bool, error) {
	// check current running task
	succeededTasks := checkRunningTask(processor, execution, db)

	// start next task
	runNextTask, err := runNextTask(processor, execution, succeededTasks, db)
	if err != nil {
		return false, err
	}
	// if start next task, return false
	if runNextTask {
		return false, nil
	}

	// if all task not completed, return false
	uncompletedTasks := processor.TaskExecutionDao.FindUncompletedByWorkflowId(execution.ID, db)
	if len(uncompletedTasks) > 0 {
		return false, nil
	}

	// if all task completed, end workflow
	endWorkflow(execution, processor, db)
	return true, nil
}

func checkRunningTask(processor *RunningStateProcessor, execution *model.WorkflowExecution, db *gorm.DB) []*model.TaskExecution {
	// load tasks
	tes := processor.TaskExecutionDao.FindUncompletedByWorkflowId(execution.ID, db)
	succeededTasks := make([]*model.TaskExecution, 0)
	for _, te := range tes {
		changed, err := te.Poll(processor.KubeClient)
		if err != nil {
			log.Logger.Error(err)
			continue
		}
		if changed {
			if te.Status == model.TaskSuccess {
				succeededTasks = append(succeededTasks, te)
			}
			processor.TaskExecutionDao.Update(te, db)
		}
	}
	return succeededTasks
}

func runNextTask(processor *RunningStateProcessor, execution *model.WorkflowExecution, succeededTasks []*model.TaskExecution, db *gorm.DB) (bool, error) {
	// start next task
	runNextTask := false
	jobDef := execution.GetJobDef()
	for _, te := range succeededTasks {
		next := jobDef.GetNextTask(te)
		if next == nil {
			continue
		}
		runNextTask = true

		log.Logger.Info(fmt.Sprintf("Run next task. current-exec-name=%s next-name=%s", te.ExecutionName, next.GetName()))
		executor, err := processor.TaskExecutorFactory.GetTaskExecutor(next)
		if err != nil {
			return runNextTask, err
		}
		err = executor.Execute(execution, db, te.Output)
		if err != nil {
			return runNextTask, errors.New(fmt.Sprintf("Failed to request task. ExecutionName=%s cause=%s", next.GetName(), err))
		}
	}
	return runNextTask, nil
}

func endWorkflow(execution *model.WorkflowExecution, processor *RunningStateProcessor, db *gorm.DB) {
	completedTasks := processor.TaskExecutionDao.FindCompletedByWorkflowId(execution.ID, db)
	hasFailedTask := false

	// failed task exists?
	for _, te := range completedTasks {
		if te.Status == model.TaskFailed {
			hasFailedTask = true
			break
		}
	}

	endedAt := time.Now()
	execution.EndedAt = &endedAt
	if hasFailedTask {
		execution.Status = model.WfFailed
	} else {
		execution.Status = model.WfSuccess
	}
	processor.WorkflowExecutionDao.Update(execution, db)
}
