package workflowstate

import (
	"testing"

	"github.com/Attsun1031/jobnetes/model"
)

func TestWorkflowExecutionStateProcessorRegistryImpl_GetProcessor(t *testing.T) {
	scheduleState := &ScheduledStateProcessor{}
	target := &RegistryImpl{
		ScheduleState: scheduleState,
	}

	result, err := target.GetProcessor(&model.WorkflowExecution{Status: model.WfScheduled})
	if err != nil {
		t.Error(err)
	} else if result != scheduleState {
		t.Error("Unexpected scheduled state")
	}
}
