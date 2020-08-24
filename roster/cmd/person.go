package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

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
		if _, err := sdk.ChannelExecute("DelPerson", personId); err != nil {
			fmt.Println("Delete person fail", err)
		}
	},
}

var personModifyCmd = &cobra.Command{
	Use:   "modify",
	Short: "Modify person info",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Modify person with id:", personId)
		bytes, err := sdk.ChannelQuery("GetPersonById", personId)
		if err != nil {
			fmt.Println("Get Person fail")
			return
		}
		person := &datatype.Person{}
		if err := json.Unmarshal(bytes, person); err != nil {
			fmt.Println("Unmarshal Person fail")
			return
		}
		_, err = sdk.ChannelExecute(
			"ModifyPerson",
			personId,
			person.Name,
			strconv.FormatUint(uint64(person.Age)+1, 10),
			strconv.FormatUint(uint64(person.Gender), 10),
			strconv.FormatInt(person.Birth.Unix(), 10),
			person.BirthPlace,
		)
		if err != nil {
			log.Println("Modify Person fail", err)
		}
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
			return
		}
		var persons []datatype.Person
		if err := json.Unmarshal(res, &persons); err != nil {
			fmt.Println("Search Person failed:", err)
			return
		}
		if len(persons) == 0 {
			fmt.Println("Person not found.")
			return
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
