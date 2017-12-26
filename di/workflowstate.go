package di

import (
	"github.com/Attsun1031/jobnetes/manager/workflowstate"
	"k8s.io/client-go/kubernetes"
)

func InjectWorkflowExecutionStateProcessorFactory(kubeClient kubernetes.Interface) workflowstate.Registry {
	return &workflowstate.RegistryImpl{
		ScheduleState: &workflowstate.ScheduledStateProcessor{
			WorkflowExecutionDao: InjectWorkflowExecutionDao(),
			TaskExecutorFactory:  InjectTaskExecutorFactory(kubeClient),
		},
	}
}
