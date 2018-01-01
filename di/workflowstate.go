package di

import (
	"github.com/Attsun1031/jobnetes/manager/workflowstate"
	"k8s.io/client-go/kubernetes"
)

func InjectWorkflowExecutionStateProcessorFactory(kubeClient kubernetes.Interface) workflowstate.Registry {
	workflowExecutionDao := InjectWorkflowExecutionDao()
	taskExecutionDao := InjectTaskExecutionDao()
	taskExecutorFactory := InjectTaskExecutorFactory(kubeClient)
	return &workflowstate.RegistryImpl{
		ScheduleState: &workflowstate.ScheduledStateProcessor{
			WorkflowExecutionDao: workflowExecutionDao,
			TaskExecutorFactory:  taskExecutorFactory,
		},
		RunningState: &workflowstate.RunningStateProcessor{
			WorkflowExecutionDao: workflowExecutionDao,
			TaskExecutionDao:     taskExecutionDao,
			TaskExecutorFactory:  taskExecutorFactory,
			KubeClient:           kubeClient,
		},
	}
}
