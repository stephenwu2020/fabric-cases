package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/stephenwu2020/fabric-cases/roster/sdk"

	"github.com/stephenwu2020/fabric-cases/roster/formater"

	"github.com/stephenwu2020/fabric-cases/chaincode/roster/datatype"

	"github.com/spf13/cobra"
)

var personCmd = &cobra.Command{
	Use:   "person",
	Short: "Add person, delete person and modify person",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var personAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add person",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add person with name:", personName)
		if _, err := sdk.ChannelExecute("AddPerson", personName); err != nil {
			fmt.Println("Add person error:", err)
		}
	},
}

var personDelCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete person",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Delete person with id:", personId)
	},
}

var personModifyCmd = &cobra.Command{
	Use:   "modify",
	Short: "Modify person info",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Modify person with id:", personId)
	},
}

var personSearchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search person",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Search person by name:", personName)
		res, err := sdk.ChannelQuery("SearchPerson", personName)
		if err != nil {
			log.Println("Search person error:", err)
			os.Exit(1)
		}
		var persons []datatype.Person
		if err := json.Unmarshal(res, &persons); err != nil {
			fmt.Println("Search Person failed:", err)
			os.Exit(1)
		}
		for i, person := range persons {
			formater.PrintPerson(i, &person)
		}
	},
}

func init() {
	rootCmd.AddCommand(personCmd)
	personCmd.AddCommand(personAddCmd)
	personCmd.AddCommand(personDelCmd)
	personCmd.AddCommand(personModifyCmd)
	personCmd.AddCommand(personSearchCmd)

	personAddCmd.Flags().StringVarP(&personName, "name", "n", "", "person name")
	personAddCmd.MarkFlagRequired("name")

	personDelCmd.Flags().StringVarP(&personId, "id", "", "", "person id")
	personDelCmd.MarkFlagRequired("id")

	personModifyCmd.Flags().StringVarP(&personId, "id", "", "", "person id")
	personModifyCmd.MarkFlagRequired("id")

	personSearchCmd.Flags().StringVarP(&personName, "name", "n", "", "person name")
	personSearchCmd.MarkFlagRequired("name")
}