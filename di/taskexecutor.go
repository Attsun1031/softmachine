package di

import (
	"github.com/Attsun1031/jobnetes/manager/taskexecutor"
	"k8s.io/client-go/kubernetes"
)

func InjectTaskExecutorFactory(kubeClient kubernetes.Interface) taskexecutor.Factory {
	return &taskexecutor.FactoryImpl{
		TaskExecutionDao: InjectTaskExecutionDao(),
		KubeClient:       kubeClient,
	}
}
