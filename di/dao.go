package di

import (
	"github.com/Attsun1031/jobnetes/dao"
)

func InjectWorkflowExecutionDao() dao.WorkflowExecutionDao {
	return &dao.WorkflowExecutionDaoImpl{}
}

func InjectTaskExecutionDao() dao.TaskExecutionDao {
	return &dao.TaskExecutionDaoImpl{}
}
