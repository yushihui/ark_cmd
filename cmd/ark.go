package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yushihui/ark_cmd/ark"
)

var from, end string

var fundCmd = &cobra.Command{
	Use:   "fund",
	Short: "View fund activity",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 1 {
			if !isValidFund(args[0]) {
				return errors.New("invalid fund name, please use arkf, arkk, arkq, arkw, arkg, izrl, prnt")
			}
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		fund := strings.ToUpper(args[0])
		end, _ := cmd.Flags().GetString("end")
		from, _ := cmd.Flags().GetString("start")
		if fund == "ALL" {
			ark.AllFundsActivity(from, end)
		} else {
			ark.FundActivity(fund, from, end)
		}

		return nil
	},
}

func isValidFund(name string) bool {
	_, ok := ark.ArkFunds[strings.ToUpper(name)]
	if ok {
		return ok
	}
	return strings.ToUpper(name) == "ALL"
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "View all ark funds",
	RunE: func(cmd *cobra.Command, args []string) error {
		for k, v := range ark.ArkFunds {
			fmt.Printf("%-10s %-30s\n", k, v)
		}
		return nil
	},
}

func init() {
	fundCmd.Flags().StringVarP(&from, "start", "s", "", "start date fotmat 2021-01-05")
	fundCmd.Flags().StringVarP(&end, "end", "e", "", "end date fotmat 2021-01-06")
	fundCmd.MarkFlagRequired("from")
	fundCmd.MarkFlagRequired("end")
	rootCmd.AddCommand(fundCmd)
	rootCmd.AddCommand(listCmd)
}
