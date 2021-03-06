package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var appName = "hf-simple-app"

var rootCmd = &cobra.Command{
	Use:   appName,
	Short: "Simple abstore cli app",
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
