package main

import (
	"fmt"

	"github.com/Attsun1031/jobnetes/dao/db"
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
	d := db.Connect(config.JobnetesConfig.DbConfig)
	defer d.Close()
	d.SetLogger(log.Logger)
	webAdminConfig := config.JobnetesConfig.WebAdminConfig

	// Disable Console Color
	// gin.DisableConsoleColor()

	// Creates a gin router with default middleware: logger and recovery (crash-free) middleware
	gin.Logger()
	router := gin.New()
	router.Use(middleware.Logger, gin.Recovery())
	router.Use(gin.BasicAuth(gin.Accounts{
		webAdminConfig.Username: webAdminConfig.Password,
	}))

	apiV1 := router.Group("/api/v1")
	{
		apiV1.GET("/workflow", api.WorkflowApi)
		apiV1.GET("/workflow/exec", func(c *gin.Context) {
			log.Logger.Info("workflow executions")
		})
		apiV1.GET("/task", func(c *gin.Context) {
			log.Logger.Info("tasks")
		})
	}

	// By default it serves on :8080 unless a PORT environment variable was defined.
	router.Run(fmt.Sprintf(":%v", webAdminConfig.Port))
}
