package di

import (
	"github.com/Attsun1031/jobnetes/manager/workflowstate"
)

func InjectWorkflowExecutionStateProcessorFactory() workflowstate.WorkflowExecutionStateProcessorRegistry {
	return &workflowstate.WorkflowExecutionStateProcessorRegistryImpl{
		ScheduleState: &workflowstate.ScheduledStateProcessor{
			WorkflowExecutionDao: InjectWorkflowExecutionDao(),
		},
	}
}
