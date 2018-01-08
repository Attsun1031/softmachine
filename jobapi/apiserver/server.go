package apiserver

import (
	"github.com/Attsun1031/jobnetes/dao"
	"github.com/Attsun1031/jobnetes/dao/db"
	"github.com/Attsun1031/jobnetes/utils/config"
	"github.com/Attsun1031/jobnetes/utils/log"
	"github.com/jinzhu/gorm"
)

type JobApiServerImpl struct {
	WorkflowDao          dao.WorkflowDao
	WorkflowExecutionDao dao.WorkflowExecutionDao
	TaskExecutionDao     dao.TaskExecutionDao
}

func (s *JobApiServerImpl) connect() *gorm.DB {
	d := db.Connect(config.JobnetesConfig.DbConfig)
	d.SetLogger(log.Logger)
	return d
}
