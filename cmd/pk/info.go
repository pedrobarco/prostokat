package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Display information about the current pk environment",
	Long:  "Display information about the current pk environment",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pk info")
	},
}
