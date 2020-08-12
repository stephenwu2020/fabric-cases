package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/stephenwu2020/fabric-cases/cli-app/cli/sdk"
)

var (
	name string
)

func init() {
	rootCmd.AddCommand(queryCmd)

	queryCmd.Flags().StringVarP(&name, "name", "n", "", "user name: a or b")
	queryCmd.MarkFlagRequired("name")
}

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Query someone's balance",
	Run: func(cmd *cobra.Command, args []string) {
		sdk.Init()
		bytes, err := sdk.ChannelQuery("query", name)
		if err != nil {
			fmt.Println("Query failed", err)
			os.Exit(1)
		}
		var balance int
		if err := json.Unmarshal(bytes, &balance); err != nil {
			fmt.Println("Unmarshal balance failed", err)
			os.Exit(1)
		}
		fmt.Println("Balance is:", balance)
	},
}
