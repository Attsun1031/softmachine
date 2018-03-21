package api

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"

	"time"

	"github.com/Attsun1031/jobnetes/dao"
	"github.com/Attsun1031/jobnetes/dao/db"
	"github.com/Attsun1031/jobnetes/utils/config"
	"github.com/Attsun1031/jobnetes/utils/log"
	"github.com/gin-gonic/gin"
	v1meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type TaskExecutionApi struct {
	TaskExecutionDao dao.TaskExecutionDao
	Client           kubernetes.Interface
}

type PodResponse struct {
	PodName    string
	StartTime  time.Time
	Containers []string
}

func (api TaskExecutionApi) GetPods(c *gin.Context) {
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

	ns := config.JobnetesConfig.KubernetesConfig.JobNamespace
	podListResult, err := api.Client.
		CoreV1().
		Pods(ns).
		List(v1meta.ListOptions{LabelSelector: fmt.Sprintf("job-name=%v", te.ExecutionName)})
	if err != nil {
		panic(err)
	}

	pods := podListResult.Items

	// sort by start time desc
	sort.SliceStable(pods, func(i, j int) bool {
		return pods[i].Status.StartTime.Before(pods[j].Status.StartTime)
	})

	podResponses := make([]*PodResponse, len(pods))
	for i, p := range pods {
		resultContainers := make([]string, len(p.Spec.Containers))
		for i, c := range p.Spec.Containers {
			resultContainers[i] = c.Name
		}
		podResponses[i] = &PodResponse{
			PodName:    p.Name,
			StartTime:  p.Status.StartTime.Time,
			Containers: resultContainers,
		}
	}
	c.JSON(http.StatusOK, gin.H{"items": podResponses})
}
