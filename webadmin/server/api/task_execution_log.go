package api

import (
	"net/http"

	"time"

	"strconv"

	"github.com/Attsun1031/jobnetes/dao"
	"github.com/Attsun1031/jobnetes/dao/db"
	"github.com/Attsun1031/jobnetes/utils/config"
	"github.com/Attsun1031/jobnetes/utils/log"
	"github.com/gin-gonic/gin"
	"k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

type TaskExecutionLogApi struct {
	Client           kubernetes.Interface
	TaskExecutionDao dao.TaskExecutionDao
}

type PodLogResponse struct {
	Name      string
	StartTime time.Time
	Log       string
}

func (api TaskExecutionLogApi) Get(c *gin.Context) {
	d := db.Connect(config.JobnetesConfig.DbConfig, log.Logger)
	defer d.Close()

	id64, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Logger.Warnf("Failed to parse request. error=%v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	id := uint(id64)

	te, err := api.TaskExecutionDao.FindById(id, d)
	if err != nil {
		panic(err)
	}
	if te == nil {
		log.Logger.Warnf("Task not found. id=%v", id)
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	pod := c.Param("pod")
	container := c.Param("container")

	ns := config.JobnetesConfig.KubernetesConfig.JobNamespace
	body, err := api.Client.
		CoreV1().
		Pods(ns).
		GetLogs(pod, &v1.PodLogOptions{Container: container}).
		Do().
		Raw()

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"item": string(body)})
}
