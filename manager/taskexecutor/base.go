package taskexecutor

import (
	"github.com/Attsun1031/jobnetes/model"
	"github.com/jinzhu/gorm"
)

type TaskExecutor interface {
	Execute(we *model.WorkflowExecution, db *gorm.DB, input string, parentId uint, prevId uint) (*model.TaskExecution, error)
}
