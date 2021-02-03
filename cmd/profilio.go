package cmd

import (
	"errors"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yushihui/ark_cmd/ark"
)

var profilioCmd = &cobra.Command{
	Use:   "profolio",
	Short: "View fund profolio",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 1 {
			if !isValidFund(args[0]) {
				return errors.New("invalid fund name, please try arkf, arkk, arkq, arkw, arkg, izrl, prnt")
			}

		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		ark.Profilio(strings.ToUpper(args[0]))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(profilioCmd)
}
