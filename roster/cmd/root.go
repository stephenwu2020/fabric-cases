package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var appName = "roster"

var rootCmd = &cobra.Command{
	Use:   appName,
	Short: "Roster records person info on Hyperledger Fabric blockchain",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
