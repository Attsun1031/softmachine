package workflowstate

import (
	"github.com/Attsun1031/jobnetes/model"
	"github.com/jinzhu/gorm"
)

type StateProcessor interface {
	// Change execution state to Next.
	// return true if change executed.
	ToNextState(*model.WorkflowExecution, *gorm.DB) (bool, error)
}
