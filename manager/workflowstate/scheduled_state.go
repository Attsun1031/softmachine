package workflowstate

import (
	"time"

	"errors"
	"fmt"

	"github.com/Attsun1031/jobnetes/dao"
	"github.com/Attsun1031/jobnetes/manager/taskexecutor"
	"github.com/Attsun1031/jobnetes/model"
	"github.com/Attsun1031/jobnetes/utils/log"
	"github.com/jinzhu/gorm"
)

type ScheduledStateProcessor struct {
	WorkflowExecutionDao dao.WorkflowExecutionDao
	TaskExecutorFactory  taskexecutor.Factory
}

func (scheduledState *ScheduledStateProcessor) ToNextState(execution *model.WorkflowExecution, db *gorm.DB) (bool, error) {
	log.Logger.Info("Scheduled state.")

	// Read job definition
	jobDef := execution.GetJobDef()

	// Start first job
	task := jobDef.GetStartTask()
	executor, err := scheduledState.TaskExecutorFactory.GetTaskExecutor(task)
	if err != nil {
		return false, err
	}
	err = executor.Execute(execution, db, execution.Input, 0)
	if err != nil {
		return false, errors.New(fmt.Sprintf("Failed to request task. ExecutionName=%s cause=%s", task.GetName(), err))
	}

	// Change state
	startTime := time.Now()
	execution.Status = model.WfRunning
	execution.StartedAt = &startTime
	err = scheduledState.WorkflowExecutionDao.Update(execution, db)
	if err != nil {
		return false, err
	}

	return true, nil
}
