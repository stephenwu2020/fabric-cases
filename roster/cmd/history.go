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
		fmt.Println("Show history with id:", argHistoryId)
	},
}

var historyAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add record to history",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add record into history:", argRecordContent, argRecordComment)
	},
}

var historyModifyCmd = &cobra.Command{
	Use:   "modify",
	Short: "Modify record of history",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Modify record of history", argRecordId, argRecordContent, argRecordComment)
	},
}

func init() {
	rootCmd.AddCommand(historyCmd)
	historyCmd.AddCommand(historyAddCmd)
	historyCmd.AddCommand(historyModifyCmd)
	historyCmd.AddCommand(historyShowCmd)

	historyAddCmd.Flags().StringVarP(&argRecordContent, "content", "c", "", "record's content")
	historyAddCmd.Flags().StringVarP(&argRecordComment, "comment", "C", "", "record's comment")
	historyAddCmd.MarkFlagRequired("content")

	historyModifyCmd.Flags().StringVarP(&argHistoryId, "id", "", "", "record's id")
	historyModifyCmd.Flags().StringVarP(&argRecordContent, "content", "c", "", "record's content")
	historyModifyCmd.Flags().StringVarP(&argRecordComment, "comment", "C", "", "record's comment")
	historyModifyCmd.MarkFlagRequired("id")

	historyShowCmd.Flags().StringVarP(&argHistoryId, "id", "", "", "record's id")
	historyShowCmd.MarkFlagRequired("id")
}
