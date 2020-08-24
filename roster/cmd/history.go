package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "Person history operations",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var historyShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show history records",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Show history with id:", historyId)
	},
}

var historyAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add record to history",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add record into history:", recordContent, recordComment)
	},
}

var historyModifyCmd = &cobra.Command{
	Use:   "modify",
	Short: "Modify record of history",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Modify record of history", recordId, recordContent, recordComment)
	},
}

func init() {
	rootCmd.AddCommand(historyCmd)
	historyCmd.AddCommand(historyAddCmd)
	historyCmd.AddCommand(historyModifyCmd)
	historyCmd.AddCommand(historyShowCmd)

	historyAddCmd.Flags().StringVarP(&recordContent, "content", "c", "", "record's content")
	historyAddCmd.Flags().StringVarP(&recordComment, "comment", "C", "", "record's comment")
	historyAddCmd.MarkFlagRequired("content")

	historyModifyCmd.Flags().StringVarP(&historyId, "id", "", "", "record's id")
	historyModifyCmd.Flags().StringVarP(&recordContent, "content", "c", "", "record's content")
	historyModifyCmd.Flags().StringVarP(&recordComment, "comment", "C", "", "record's comment")
	historyModifyCmd.MarkFlagRequired("id")

	historyShowCmd.Flags().StringVarP(&historyId, "id", "", "", "record's id")
	historyShowCmd.MarkFlagRequired("id")
}
