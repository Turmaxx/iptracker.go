package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "routescout",
		Short: "routescout CLI app.",
		Long:  `routescout CLI app.`,
	}
)

func Execute() error {
	return rootCmd.Execute()
}
