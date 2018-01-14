package manager

import (
	"time"

	"github.com/Attsun1031/jobnetes/dao"
	"github.com/Attsun1031/jobnetes/manager/workflowstate"
	"github.com/Attsun1031/jobnetes/utils/log"
	"github.com/jinzhu/gorm"
)

const pollIntervalSec = 10 * time.Second

type WorkflowManagerMain struct {
	Db                                 *gorm.DB
	WorkflowExecutionDao               dao.WorkflowExecutionDao
	WorkflowExecutionProcessorRegistry workflowstate.Registry
}

func (manager *WorkflowManagerMain) Run() {
	for {
		log.Logger.Info("Looping")
		manager.processWorkflowState()
		time.Sleep(pollIntervalSec)
	}
}

func (manager *WorkflowManagerMain) processWorkflowState() {
	// Load running workflows
	log.Logger.Info("Running workflow state manager")

	uncompletedExecs, err := manager.WorkflowExecutionDao.FindUncompletedWorkflowExecs(manager.Db)
	if err != nil {
		log.Logger.Fatalf("Failed to load workflow executions. err=%v", err)
		return
	}

	for _, exec := range uncompletedExecs {
		log.Logger.Infof("Process WorkflowExecution: id=%d", exec.ID)

		// get state object
		prevState := exec.Status

		stateProcessor, err := manager.WorkflowExecutionProcessorRegistry.GetProcessor(exec)
		if err != nil {
			log.Logger.Error("Failed to get Workflow Execution processor. id=%d cause=%s", exec.ID, err)
			continue
		}

		tx := manager.Db.Begin()
		stateChanged, err := stateProcessor.ToNextState(exec, tx)
		if err != nil {
			tx.Rollback()
			log.Logger.Errorf("Failed to change state for wid=%d cause=%s", exec.ID, err)
			continue
		}
		tx.Commit()

		newState := exec.Status
		if stateChanged {
			log.Logger.Infof("Workflow state changed. id=%d prev=%d to=%d", exec.ID, prevState, newState)
		} else if err != nil {
			log.Logger.Errorf("Failed to change state. id=%d cause='%s'", exec.ID, err)
		}
	}
}
