package app

import (
	"fmt"

	"github.com/Attsun1031/jobnetes/dao/db"
	"github.com/Attsun1031/jobnetes/di"
	"github.com/Attsun1031/jobnetes/kubernetes"
	"github.com/Attsun1031/jobnetes/manager"
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
	fmt.Println("jobnetesmanager run")
	manager.InitConfig()
	d := db.Connect(manager.ManagerConfig.DbConfig)
	kubeClient := kubernetes.GetClient(manager.ManagerConfig.KubeConfig)
	mgr := &manager.WorkflowManagerMain{
		Db:                                 d,
		WorkflowExecutionDao:               di.InjectWorkflowExecutionDao(),
		WorkflowExecutionProcessorRegistry: di.InjectWorkflowExecutionStateProcessorFactory(kubeClient),
	}
	mgr.Run()
}
