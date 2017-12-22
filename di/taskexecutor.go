package di

import (
	"github.com/Attsun1031/jobnetes/manager/taskexecutor"
)

func InjectTaskExecutorFactory() taskexecutor.Factory {
	return &taskexecutor.FactoryImpl{}
}
