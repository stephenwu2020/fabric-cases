package cmd

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/stephenwu2020/fabric-cases/chaincode/roster/datatype"
	"github.com/stephenwu2020/fabric-cases/roster/formater"
	"github.com/stephenwu2020/fabric-cases/roster/sdk"
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
		fmt.Println("Show history with person id:", argPersonId)
		history, err := getHistoryByPersonId(argPersonId)
		if err != nil {
			fmt.Println(err)
			return
		}
		for i, r := range history.Records {
			formater.PrintRecord(i, &r)
		}
	},
}

var historyAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add record to history",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add record into history:", argRecordContent, argRecordComment)
		history, err := getHistoryByPersonId(argPersonId)
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = sdk.ChannelExecute("AddHistory", history.Id, argRecordContent, argRecordComment)
		if err != nil {
			fmt.Println("Add history fail", err)
			return
		}
		fmt.Println("Add history record success.")
	},
}

var historyModifyCmd = &cobra.Command{
	Use:   "modify",
	Short: "Modify record of history",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Modify record of history", argRecordId, argRecordContent, argRecordComment)
		history, err := getHistoryByPersonId(argPersonId)
		if err != nil {
			fmt.Println(err)
			return
		}
		rid, err := strconv.Atoi(argRecordId)
		if err != nil {
			fmt.Println("Parse record id fail")
			return
		}
		if rid >= len(history.Records) {
			fmt.Println("Record id out of range")
			return
		}
		old := history.Records[rid]
		fcontent := old.Content
		if argRecordContent != "" {
			fcontent = argRecordContent
		}
		fcomment := old.Comment
		if argRecordComment != "" {
			fcomment = argRecordComment
		}

		_, err = sdk.ChannelExecute(
			"ModifyHistory",
			history.Id,
			argRecordId,
			fcontent,
			fcomment,
		)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Modify record success.")
	},
}

func init() {
	rootCmd.AddCommand(historyCmd)
	historyCmd.AddCommand(historyAddCmd)
	historyCmd.AddCommand(historyModifyCmd)
	historyCmd.AddCommand(historyShowCmd)

	historyAddCmd.Flags().StringVarP(&argPersonId, "id", "", "", "person id")
	historyAddCmd.Flags().StringVarP(&argRecordContent, "content", "c", "", "record's content")
	historyAddCmd.Flags().StringVarP(&argRecordComment, "comment", "C", "", "record's comment")
	historyModifyCmd.MarkFlagRequired("id")
	historyAddCmd.MarkFlagRequired("content")

	historyModifyCmd.Flags().StringVarP(&argPersonId, "id", "", "", "person's id")
	historyModifyCmd.Flags().StringVarP(&argRecordId, "rid", "", "", "record's id")
	historyModifyCmd.Flags().StringVarP(&argRecordContent, "content", "c", "", "record's content")
	historyModifyCmd.Flags().StringVarP(&argRecordComment, "comment", "C", "", "record's comment")
	historyModifyCmd.MarkFlagRequired("id")
	historyModifyCmd.MarkFlagRequired("rid")

	historyShowCmd.Flags().StringVarP(&argPersonId, "id", "", "", "person id")
	historyShowCmd.MarkFlagRequired("id")
}

func getHistoryByPersonId(personId string) (*datatype.History, error) {
	person, err := getPersonById(personId)
	if err != nil {
		return nil, errors.WithMessage(err, "Get person fail")
	}
	bytes, err := sdk.ChannelQuery("GetHistory", person.HistroyId)
	if err != nil {
		return nil, errors.WithMessage(err, "Get history fail")
	}
	var history datatype.History
	if err := json.Unmarshal(bytes, &history); err != nil {
		return nil, errors.WithMessage(err, "Marshal history fail")
	}
	return &history, nil
}
