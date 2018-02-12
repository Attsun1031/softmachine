package api

import (
	"net/http"

	"strconv"

	"github.com/Attsun1031/jobnetes/dao"
	"github.com/Attsun1031/jobnetes/dao/db"
	"github.com/Attsun1031/jobnetes/utils/config"
	"github.com/Attsun1031/jobnetes/utils/log"
	"github.com/gin-gonic/gin"
)

type WorkflowExecutionDetailApi struct {
	WorkflowDao          dao.WorkflowDao
	WorkflowExecutionDao dao.WorkflowExecutionDao
	TaskExecutionDao     dao.TaskExecutionDao
}

func (api *WorkflowExecutionDetailApi) Get(c *gin.Context) {
	id64, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Logger.Warnf("Faield to parse request. error=%v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	id := uint(id64)

	d := db.Connect(config.JobnetesConfig.DbConfig, log.Logger)
	defer d.Close()
	log.Logger.Infof("workflow execution detail with id=%v", id)
	we, err := api.WorkflowExecutionDao.FindById(id, true, d)
	if err != nil {
		panic(err)
	}
	log.Logger.Infof("tasks = %v", len(we.TaskExecutions))
	c.JSON(http.StatusOK, gin.H{"item": we})
}
