package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yushihui/ark_cmd/ark"
)

var arkFunds = map[string]string{
	"ARKK": "ARK INNOVATION ~ 20+ B",
	"ARKQ": "ARK AUTONOMOUS TECHNOLOGY & ROBOTICS ~ 2.7 B",
	"ARKW": "ARK NEXT GENERATION INTERNET ~ 6.5 B",
	"ARKG": "ARK GENOMIC REVOLUTION MULTISECTOR ~ 11+ B",
	"ARKF": "ARK FINTECH INNOVATION ~ 2.6 B",
	"PRNT": "THE 3D RINTING ~ 0.26 B",
	"IZRL": "ARK ISRAEL INNOVATIVE TECHNOLOGY ~ 0.15 B",
}

var from, end string

var viewCmd = &cobra.Command{
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
		from, _ := cmd.Flags().GetString("from")
		ark.Diff(fund, from, end)
		return nil
	},
}

func isValidFund(name string) bool {
	_, ok := arkFunds[strings.ToUpper(name)]
	return ok
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "View all ark funds",
	RunE: func(cmd *cobra.Command, args []string) error {
		for k, v := range arkFunds {
			fmt.Printf("%-10s %-30s\n", k, v)
		}

		return nil
	},
}

func init() {
	viewCmd.Flags().StringVarP(&from, "from", "f", "", "from time fotmat 2021-01-05")
	viewCmd.Flags().StringVarP(&end, "end", "e", "", "end time fotmat 2021-01-06")
	viewCmd.MarkFlagRequired("from")
	viewCmd.MarkFlagRequired("end")
	rootCmd.AddCommand(viewCmd)
	rootCmd.AddCommand(listCmd)
}
