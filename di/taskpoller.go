package di

import (
	"github.com/Attsun1031/jobnetes/manager/taskpoller"
	"k8s.io/client-go/kubernetes"
)

func InjectTaskPollerFactory(kubeClient kubernetes.Interface) taskpoller.Factory {
	return &taskpoller.FactoryImpl{
		TaskExecutionDao: InjectTaskExecutionDao(),
		KubeClient:       kubeClient,
	}
}
