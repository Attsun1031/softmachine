package manager

import (
	"time"

	"github.com/Attsun1031/jobnetes/dao"
	"github.com/Attsun1031/jobnetes/manager/workflowstate"
	"github.com/Attsun1031/jobnetes/utils/log"
	"github.com/jinzhu/gorm"
)

const pollInterval = 10 * time.Second

type WorkflowManagerMain struct {
	Db                                 *gorm.DB
	WorkflowExecutionDao               dao.WorkflowExecutionDao
	WorkflowExecutionProcessorRegistry workflowstate.Registry
}

func (manager *WorkflowManagerMain) Run() {
	for {
		log.Logger.Info("Looping")
		manager.processWorkflowState()
		time.Sleep(pollInterval)
	}
}

// TODO: retry failed workflow
func (manager *WorkflowManagerMain) processWorkflowState() {
	// Load running workflows
	log.Logger.Info("Running workflow state manager")

	uncompletedExecs, err := manager.WorkflowExecutionDao.FindUncompletedWorkflowExecs(manager.Db)
	if err != nil {
		log.Logger.Fatalf("Failed to load workflow executions. err=%v", err)
		return
	}

	for _, exec := range uncompletedExecs {
		log.Logger.Infof("Process WorkflowExecution: id=%v", exec.ID)

		// get state object
		prevState := exec.Status

		stateProcessor, err := manager.WorkflowExecutionProcessorRegistry.GetProcessor(exec)
		if err != nil {
			log.Logger.Errorf("Failed to get Workflow Execution processor. id=%v cause=%v", exec.ID, err)
			continue
		}

		// ensure a WorkflowExecution record is handled only by this thread.
		// (as a result, TaskExecutions related to the WorkflowExecution are handled only by this.)
		tx := manager.Db.Begin()
		execCurrent, err := manager.WorkflowExecutionDao.FindById(exec.ID, false, false, tx.Set("gorm:query_option", "FOR UPDATE"))
		if err != nil {
			log.Logger.Errorf("Failed to fetch and lock target workflow execution record. id=%v cause=%v", exec.ID, err)
			tx.Rollback()
			continue
		}
		if execCurrent.Status != exec.Status {
			log.Logger.Info("State changed by other process.")
			tx.Rollback()
			continue
		}

		stateChanged, err := stateProcessor.ToNextState(execCurrent, tx)
		if err != nil {
			tx.Rollback()
			log.Logger.Errorf("Failed to change state for wid=%v cause=%v", exec.ID, err)
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
