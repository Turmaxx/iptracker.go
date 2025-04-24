package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Shows version",
		Long:  "Displays the version of the application",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("v1.2")
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}
