package api

import (
	"net/http"

	"github.com/Attsun1031/jobnetes/dao"
	"github.com/Attsun1031/jobnetes/dao/db"
	"github.com/Attsun1031/jobnetes/model"
	"github.com/Attsun1031/jobnetes/utils/config"
	"github.com/Attsun1031/jobnetes/utils/log"
	"github.com/gin-gonic/gin"
)

type WorkflowExecutionApi struct {
	WorkflowDao          dao.WorkflowDao
	WorkflowExecutionDao dao.WorkflowExecutionDao
}

func (api *WorkflowExecutionApi) Get(c *gin.Context) {
	d := db.Connect(config.JobnetesConfig.DbConfig, log.Logger)
	defer d.Close()
	log.Logger.Info("workflow executions")
	execs, err := api.WorkflowExecutionDao.Find(100, 0, "ended_at desc", d)
	if err != nil {
		panic(err)
	}

	wids := make([]uint, 0)
	widSet := make(map[uint]struct{})
	for _, we := range execs {
		widSet[we.WorkflowID] = struct{}{}
	}
	idToWorkflow := make(map[uint]*model.Workflow, len(widSet))

	for id := range widSet {
		wids = append(wids, id)
	}
	workflows, err := api.WorkflowDao.FindByIds(wids, d)
	for _, w := range workflows {
		idToWorkflow[w.ID] = w
	}

	for _, we := range execs {
		we.Workflow = idToWorkflow[we.WorkflowID]
	}
	c.JSON(http.StatusOK, gin.H{"items": execs})
}
