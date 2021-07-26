package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	profilesCmd = &cobra.Command{
		Use:   "profiles",
		Short: "Manage prostokat profiles",
	}

	activateProfileCmd = &cobra.Command{
		Use:   "activate",
		Short: "Activates an existing named profile",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("activate profile")
		},
	}

	createProfileCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a new named profile",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("create profile")
		},
	}

	deleteProfileCmd = &cobra.Command{
		Use:   "delete",
		Short: "Deletes a named profile",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("delete profile")
		},
	}

	describeProfileCmd = &cobra.Command{
		Use:   "describe",
		Short: "Describes a named profile by listing its properties",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("describe profile")
		},
	}

	listProfilesCmd = &cobra.Command{
		Use:   "list",
		Short: "Lists existing named profiles",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("list profiles")
		},
	}
)

func init() {
	// profilesCmd.AddCommand(activateProfileCmd)
	// profilesCmd.AddCommand(createProfileCmd)
	// profilesCmd.AddCommand(deleteProfileCmd)
	// profilesCmd.AddCommand(describeProfileCmd)
	// profilesCmd.AddCommand(listProfilesCmd)
}
