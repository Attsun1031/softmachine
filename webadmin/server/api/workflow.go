package api

import (
	"net/http"

	"github.com/Attsun1031/jobnetes/utils/log"
	"github.com/gin-gonic/gin"
)

func WorkflowApi(c *gin.Context) {
	log.Logger.Info("workflows")
	c.JSON(http.StatusOK, gin.H{"hoge": 1})
}
