package cmd

import (
	"os/exec"

	"github.com/spf13/cobra"
)

var urlCmd = &cobra.Command{
	Use:   "url",
	Short: "url command",
	RunE: func(cmd *cobra.Command, args []string) error {
		openURL()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(urlCmd)
}

func openURL() {
	exec.Command("open", "https://www.yahoo.com").Start()
}
