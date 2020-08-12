package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/stephenwu2020/fabric-cases/cli-app/cli/sdk"
)

var (
	from  string
	to    string
	value int
)

func init() {
	rootCmd.AddCommand(transferCmd)

	transferCmd.Flags().StringVarP(&from, "from", "f", "", "transfer from who")
	transferCmd.Flags().StringVarP(&to, "to", "t", "", "transfer to who")
	transferCmd.Flags().IntVarP(&value, "value", "v", 0, "transfer amount")
}

var transferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "transfer someone's balance to another one",
	Run: func(cmd *cobra.Command, args []string) {
		sdk.Init()
		valueStr := strconv.Itoa(value)
		_, err := sdk.ChannelExecute("invoke", from, to, valueStr)
		if err != nil {
			fmt.Println("Transfer failed", err)
			os.Exit(1)
		}
		fmt.Println("Transfer success!")
	},
}
