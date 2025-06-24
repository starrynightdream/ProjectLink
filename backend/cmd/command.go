package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// commandCmd Script entry
var commandCmd = &cobra.Command{
	Use:   "command",
	Short: "Call some script",
	Long: `Defind some script for database handler,
	or for some information searching.`,
}

// versionCmd show version
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show system version",
	Long:  "show system version to the screen.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(VERSION)
	},
}

// TODO more database command

func init() {
	// all Script
	commandCmd.AddCommand(versionCmd)

	// registy command
	rootCmd.AddCommand(commandCmd)
}
