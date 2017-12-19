package workflowstate

import (
	"errors"
	"fmt"

	"github.com/Attsun1031/jobnetes/model"
)

type WorkflowExecutionStateProcessorRegistry interface {
	GetProcessor(*model.WorkflowExecution) (StateProcessor, error)
}

type WorkflowExecutionStateProcessorRegistryImpl struct {
	ScheduleState *ScheduledStateProcessor
}

func (registry *WorkflowExecutionStateProcessorRegistryImpl) GetProcessor(execution *model.WorkflowExecution) (StateProcessor, error) {
	switch execution.Status {
	case model.WfScheduled:
		return registry.ScheduleState, nil
	default:
		return nil, errors.New(fmt.Sprintf("Unknown workflow state. %d", execution.Status))
	}
}
