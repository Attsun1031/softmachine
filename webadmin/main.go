package main

import (
	"github.com/Attsun1031/jobnetes/dao/db"
	"github.com/Attsun1031/jobnetes/utils/config"
	"github.com/Attsun1031/jobnetes/utils/log"
	"github.com/gin-gonic/gin"
)

func main() {
	// Setup application config
	config.InitConfig()
	log.SetupLogger(config.JobnetesConfig.LogConfig)
	d := db.Connect(config.JobnetesConfig.DbConfig)
	d.SetLogger(log.Logger)

	// Disable Console Color
	// gin.DisableConsoleColor()

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()
	gin.Logger()

	router.GET("/test", func(c *gin.Context) {
		x := c.Query("x")
		log.Logger.Infof("x = %s\n", x)
	})

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
	// router.Run(":3000") for a hard coded port
}
