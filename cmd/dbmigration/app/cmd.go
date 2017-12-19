package app

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewDbmigrationCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "dbmigration",
		Long: `This is dbmigraiton tool for jobnetes`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("dbmigration run")
			Run()
		},
	}
	return cmd
}
