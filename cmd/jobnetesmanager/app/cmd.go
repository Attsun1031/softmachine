package app

import (
	"github.com/Attsun1031/jobnetes/dao/db"
	"github.com/Attsun1031/jobnetes/di"
	"github.com/Attsun1031/jobnetes/kubernetes"
	"github.com/Attsun1031/jobnetes/manager"
	"github.com/Attsun1031/jobnetes/utils/config"
	"github.com/Attsun1031/jobnetes/utils/log"
	"github.com/spf13/cobra"
)

func NewJobmanagerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "jobnetesmanager",
		Long: `This is jobnetesmanager for jobnetes`,
		Run:  run,
	}
	return cmd
}

func run(cmd *cobra.Command, args []string) {
	config.InitConfig()
	log.SetupLogger(config.JobnetesConfig.LogConfig)
	d := db.Connect(config.JobnetesConfig.DbConfig)
	defer d.Close()
	d.SetLogger(log.Logger)
	kubeClient := kubernetes.GetClient(config.JobnetesConfig.KubernetesConfig)
	mgr := &manager.WorkflowManagerMain{
		Db:                                 d,
		WorkflowExecutionDao:               di.InjectWorkflowExecutionDao(),
		WorkflowExecutionProcessorRegistry: di.InjectWorkflowExecutionStateProcessorFactory(kubeClient),
	}
	mgr.Run()
}
