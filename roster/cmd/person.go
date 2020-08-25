package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/pkg/errors"
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
		fmt.Println("Add person with name:", argPersonName)
		if _, err := sdk.ChannelExecute("AddPerson", argPersonName); err != nil {
			fmt.Println("Add person error:", err)
		}
	},
}

var personDelCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete person",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Delete person with id:", argPersonId)
		if _, err := sdk.ChannelExecute("DelPerson", argPersonId); err != nil {
			fmt.Println("Delete person fail", err)
		}
	},
}

var personModifyCmd = &cobra.Command{
	Use:   "modify",
	Short: "Modify person info",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Modify person with id:", argPersonId)
		person, err := getPersonById(argPersonId)
		if err != nil {
			fmt.Println("Get Person fail", err)
		}

		fname := person.Name
		if argPersonName != "" {
			fname = argPersonName
		}
		fage := person.Age
		if argAge != 0 {
			fage = argAge
		}
		fgender := person.Gender
		if argGender != 0 {
			fgender = argGender
		}
		fbirth := person.Birth.Unix()
		if argBirth != 0 {
			fbirth = argBirth
		}
		fplace := person.BirthPlace
		if argBirthPlace != "" {
			fplace = argBirthPlace
		}

		_, err = sdk.ChannelExecute(
			"ModifyPerson",
			argPersonId,
			fname,
			strconv.FormatUint(uint64(fage), 10),
			strconv.FormatUint(uint64(fgender), 10),
			strconv.FormatInt(fbirth, 10),
			fplace,
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
		fmt.Println("Search person by name:", argPersonName)
		res, err := sdk.ChannelQuery("SearchPerson", argPersonName)
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

	personAddCmd.Flags().StringVarP(&argPersonName, "name", "n", "", "person name")
	personAddCmd.MarkFlagRequired("name")

	personDelCmd.Flags().StringVarP(&argPersonId, "id", "", "", "person id")
	personDelCmd.MarkFlagRequired("id")

	personModifyCmd.Flags().StringVarP(&argPersonId, "id", "", "", "person id")
	personModifyCmd.Flags().StringVarP(&argPersonName, "name", "n", "", "person name")
	personModifyCmd.Flags().Uint8VarP(&argAge, "age", "a", 0, "age")
	personModifyCmd.Flags().Uint8VarP(&argGender, "gender", "g", 0, "gender")
	personModifyCmd.Flags().Int64VarP(&argBirth, "birth", "b", 0, "birth timestamp")
	personModifyCmd.Flags().StringVarP(&argBirthPlace, "place", "p", "", "birth place")
	personModifyCmd.MarkFlagRequired("id")

	personSearchCmd.Flags().StringVarP(&argPersonName, "name", "n", "", "person name")
	personSearchCmd.MarkFlagRequired("name")

}

func getPersonById(personId string) (*datatype.Person, error) {
	bytes, err := sdk.ChannelQuery("GetPersonById", personId)
	if err != nil {
		return nil, errors.WithMessage(err, "Get person by id fail")
	}
	person := &datatype.Person{}
	if err := json.Unmarshal(bytes, person); err != nil {
		return nil, errors.WithMessage(err, "Unmarshal person fail")
	}
	return person, nil
}
