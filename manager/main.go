package manager

import (
	"time"

	"fmt"

	"github.com/Attsun1031/jobnetes/dao"
	"github.com/Attsun1031/jobnetes/manager/workflowstate"
	"github.com/Attsun1031/jobnetes/model"
	"github.com/Attsun1031/jobnetes/utils/log"
	"github.com/jinzhu/gorm"
)

const pollIntervalSec = 10 * time.Second

type WorkflowManagerMain struct {
	Db                                 *gorm.DB
	WorkflowExecutionDao               dao.WorkflowExecutionDao
	WorkflowExecutionProcessorRegistry workflowstate.WorkflowExecutionStateProcessorRegistry
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
	//loadData()

	uncompletedExecs := manager.WorkflowExecutionDao.FindUncompletedWorkflowExecs(manager.Db)

	for _, exec := range uncompletedExecs {
		log.Logger.Info(fmt.Sprintf("Process WorkflowExecution: id=%d", exec.ID))

		// get state object
		prevState := exec.Status

		stateProcessor, err := manager.WorkflowExecutionProcessorRegistry.GetProcessor(exec)
		if err != nil {
			log.Logger.Error("Failed to get Workflow Execution processor. id=%d cause=%s", exec.ID, err)
			continue
		}

		tx := manager.Db.Begin()
		stateChanged, err := stateProcessor.ToNextState(exec, manager.Db)
		if err != nil {
			tx.Rollback()
			continue
		}
		tx.Commit()

		newState := exec.Status
		if stateChanged {
			log.Logger.Info(fmt.Sprintf("Workflow state changed. id=%d prev=%d to=%d", exec.ID, prevState, newState))
		} else if err != nil {
			log.Logger.Error(fmt.Sprintf("Failed to change state. id=%d cause='%s'", exec.ID, err))
		}
	}
}

func loadData(db *gorm.DB) {
	user := &model.User{
		Name: "jon",
	}
	db.Create(user)

	workflow := &model.Workflow{
		Name:       "sample",
		Definition: `{"x":"hoge"}`,
		User:       user,
	}
	db.Create(workflow)

	exec := &model.WorkflowExecution{
		Workflow: workflow,
		Name:     "exec1",
		Status:   model.WfScheduled,
		Input:    "{}",
		Output:   "{}",
	}
	db.Create(&exec)
}
