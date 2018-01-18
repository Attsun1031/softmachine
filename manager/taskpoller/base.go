package taskpoller

import (
	"github.com/Attsun1031/jobnetes/model"
	"github.com/jinzhu/gorm"
)

type TaskPoller interface {
	// Check if TaskExecution completed
	Poll(*model.TaskExecution, *gorm.DB) (bool, error)
}
