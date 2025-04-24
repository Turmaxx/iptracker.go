package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "iptracker",
		Short: "IPtracker CLI app.",
		Long:  `IPtracker CLI app.`,
	}
)

func Execute() error {
	return rootCmd.Execute()
}
