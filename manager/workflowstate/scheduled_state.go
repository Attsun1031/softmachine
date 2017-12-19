package workflowstate

import (
	"time"

	"errors"
	"fmt"

	"github.com/Attsun1031/jobnetes/dao"
	"github.com/Attsun1031/jobnetes/model"
	"github.com/Attsun1031/jobnetes/utils/log"
	"github.com/jinzhu/gorm"
)

type ScheduledStateProcessor struct {
	WorkflowExecutionDao dao.WorkflowExecutionDao
}

func (scheduledState *ScheduledStateProcessor) ToNextState(execution *model.WorkflowExecution, db *gorm.DB) (bool, error) {
	log.Logger.Info("Scheduled state.")

	// Read job definition
	jobDef := execution.Workflow.GetJobDef()

	// Load latest state
	executionCurrent := scheduledState.WorkflowExecutionDao.FindById(execution.ID, db.Set("gorm:query_option", "FOR UPDATE"))
	if executionCurrent.Status != execution.Status {
		log.Logger.Info("State changed by other process.")
		return false, nil
	}

	// Start first job
	task := jobDef.GetStartTask()
	err := task.RequestTask(executionCurrent)
	if err != nil {
		return false, errors.New(fmt.Sprintf("Failed to request task. Name=%s cause=%s", task.GetName(), err))
	}

	// Change state
	startTime := time.Now()
	executionCurrent.Status = model.WfRunning
	executionCurrent.StartedAt = &startTime
	scheduledState.WorkflowExecutionDao.Update(executionCurrent, db)

	return true, nil
}
