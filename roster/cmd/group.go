package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "Mark person with group tag",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var groupListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all group tags",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("List all group tags")
	},
}

var groupAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a group tag",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add a group tag:", groupTag)
	},
}

var groupDelCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete a group tag",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Delete a group tag:", groupTag)
	},
}

var groupAssignCmd = &cobra.Command{
	Use:   "assign",
	Short: "Assign a group tag to a person",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Assign a group tag: %s to person: %s\n", groupTag, personId)
	},
}

var groupRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a group tag of person",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Remove a group tag: %s of person: %s\n", groupTag, personId)
	},
}

func init() {
	rootCmd.AddCommand(groupCmd)
	groupCmd.AddCommand(groupListCmd)
	groupCmd.AddCommand(groupAddCmd)
	groupCmd.AddCommand(groupDelCmd)
	groupCmd.AddCommand(groupAssignCmd)
	groupCmd.AddCommand(groupRemoveCmd)

	groupAddCmd.Flags().StringVarP(&groupTag, "tag", "t", "", "group tag")
	groupAddCmd.MarkFlagRequired("tag")

	groupDelCmd.Flags().StringVarP(&groupTag, "tag", "t", "", "group tag")
	groupDelCmd.MarkFlagRequired("tag")

	groupAssignCmd.Flags().StringVarP(&groupTag, "tag", "t", "", "group tag")
	groupAssignCmd.Flags().StringVarP(&personId, "id", "", "", "person id")
	groupAssignCmd.MarkFlagRequired("tag")
	groupAssignCmd.MarkFlagRequired("id")

	groupRemoveCmd.Flags().StringVarP(&groupTag, "tag", "t", "", "group tag")
	groupRemoveCmd.Flags().StringVarP(&personId, "id", "", "", "person id")
	groupRemoveCmd.MarkFlagRequired("tag")
	groupRemoveCmd.MarkFlagRequired("id")
}
