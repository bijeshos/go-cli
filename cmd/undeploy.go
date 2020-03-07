package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(undeployCmd)

	undeployCmd.AddCommand(undeployWebCmd)
	undeployCmd.AddCommand(undeployDatabaseCmd)

}

var undeployCmd = &cobra.Command{
	Use:     "undeploy",
	Aliases: []string{"undep", "undepl"},
	Short:   "undeploy apps",
	Long:    `undeploy apps`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing undeploy")
	},
}

var undeployWebCmd = &cobra.Command{
	Use:   "web",
	Short: "undeploy web",
	Long:  `undeploy web`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing undeploy:web")
	},
}

var undeployDatabaseCmd = &cobra.Command{
	Use:   "database",
	Short: "undeploy database",
	Long:  `undeploy database`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing undeploy:database")
	},
}
