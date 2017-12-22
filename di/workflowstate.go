package di

import (
	"github.com/Attsun1031/jobnetes/manager/workflowstate"
)

func InjectWorkflowExecutionStateProcessorFactory() workflowstate.Registry {
	return &workflowstate.RegistryImpl{
		ScheduleState: &workflowstate.ScheduledStateProcessor{
			WorkflowExecutionDao: InjectWorkflowExecutionDao(),
			TaskExecutorFactory:  InjectTaskExecutorFactory(),
		},
	}
}
