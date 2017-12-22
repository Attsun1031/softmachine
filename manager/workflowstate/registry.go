package workflowstate

import (
	"errors"
	"fmt"

	"github.com/Attsun1031/jobnetes/model"
)

type Registry interface {
	GetProcessor(*model.WorkflowExecution) (StateProcessor, error)
}

type RegistryImpl struct {
	ScheduleState *ScheduledStateProcessor
}

func (registry *RegistryImpl) GetProcessor(execution *model.WorkflowExecution) (StateProcessor, error) {
	switch execution.Status {
	case model.WfScheduled:
		return registry.ScheduleState, nil
	default:
		return nil, errors.New(fmt.Sprintf("Unknown workflow state. %d", execution.Status))
	}
}
