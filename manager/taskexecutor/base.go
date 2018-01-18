package taskexecutor

import (
	"github.com/Attsun1031/jobnetes/model"
	"github.com/jinzhu/gorm"
)

type TaskExecutor interface {
	Execute(*model.WorkflowExecution, *gorm.DB, string, uint) error
}
