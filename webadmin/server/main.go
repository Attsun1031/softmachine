package main

import (
	"fmt"

	"github.com/Attsun1031/jobnetes/di"
	"github.com/Attsun1031/jobnetes/kubernetes"
	"github.com/Attsun1031/jobnetes/utils/config"
	"github.com/Attsun1031/jobnetes/utils/log"
	"github.com/Attsun1031/jobnetes/webadmin/server/api"
	"github.com/Attsun1031/jobnetes/webadmin/server/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	// Setup application config
	config.InitConfig()
	log.SetupLogger(config.JobnetesConfig.LogConfig)
	webAdminConfig := config.JobnetesConfig.WebAdminConfig

	// Disable Console Color
	// gin.DisableConsoleColor()

	// Creates a gin router with default middleware: logger and recovery (crash-free) middleware
	gin.Logger()
	router := gin.New()
	router.Use(middleware.Logger, gin.Recovery())
	//router.Use(gin.BasicAuth(gin.Accounts{
	//	webAdminConfig.Username: webAdminConfig.Password,
	//}))

	kubeClient := kubernetes.GetClient(config.JobnetesConfig.KubernetesConfig)
	apiV1 := router.Group("/api/v1")
	{
		apiV1.GET("/workflow", api.WorkflowApi)
		apiV1.GET("/workflow/execution", (&api.WorkflowExecutionApi{
			WorkflowDao:          di.InjectWorkflowDao(),
			WorkflowExecutionDao: di.InjectWorkflowExecutionDao(),
		}).Get)
		apiV1.GET("/workflow/execution/:id", (&api.WorkflowExecutionDetailApi{
			WorkflowDao:          di.InjectWorkflowDao(),
			WorkflowExecutionDao: di.InjectWorkflowExecutionDao(),
			TaskExecutionDao:     di.InjectTaskExecutionDao(),
		}).Get)
		apiV1.GET("/task/:id/pod", (&api.TaskExecutionApi{
			TaskExecutionDao: di.InjectTaskExecutionDao(),
			Client:           kubeClient,
		}).GetPods)
		apiV1.GET("/task/:id/pod/:pod/:container/log", (&api.TaskExecutionLogApi{
			TaskExecutionDao: di.InjectTaskExecutionDao(),
			Client:           kubeClient,
		}).Get)
	}

	// By default it serves on :8080 unless a PORT environment variable was defined.
	router.Run(fmt.Sprintf(":%v", webAdminConfig.Port))
}
