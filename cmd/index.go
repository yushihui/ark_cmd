package cmd

import (
	"encoding/gob"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yushihui/ark_cmd/ark"
)

var indexCmd = &cobra.Command{
	Use:   "index",
	Short: "build index",
	RunE: func(cmd *cobra.Command, args []string) error {
		BuildIndex()
		return nil
	},
}

var tickerCmd = &cobra.Command{
	Use:   "ticker",
	Short: "ticker activity",
	RunE: func(cmd *cobra.Command, args []string) error {
		searchActivity(args[0])
		return nil
	},
}

// SearchActivity by ticker
func searchActivity(ticker string) {
	if activitiesMap, err := decode(); err == nil {
		ticker = strings.ToUpper(ticker)
		if activities, ok := activitiesMap[ticker]; ok {
			printActivities(*activities)
		}
	} else {
		fmt.Println("NO activities")
	}
}

func printActivities(activities []ark.Security) {
	fmt.Printf("%-10s %10s %18s %18s %20s\n", "Date", "Ticker", "Shares", "Value", "Fund")
	for _, a := range activities {
		fmt.Printf("%4d-%02d-%02d %10s %18s %18s %20s\n", a.TradDate.Year(), a.TradDate.Month(), a.TradDate.Day(), a.Ticker, ark.PrettyNumber(a.Delta), ark.PrettyNumber(a.Delta*a.Price), a.Fund)
	}
}

// BuildIndex build index
func BuildIndex() {
	activities := make(map[string]*[]ark.Security)
	for k := range arkFund {
		ark.ParseFundActivity(k, activities)
	}

	dataFile, err := os.Create("data/index.gob")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer dataFile.Close()
	// serialize the data
	dataEncoder := gob.NewEncoder(dataFile)
	dataEncoder.Encode(activities)
	fmt.Println("Index build done")
}

// Decode data from index
func decode() (map[string]*[]ark.Security, error) {
	activities := make(map[string]*[]ark.Security)
	// open data file
	dataFile, err := os.Open("data/index.gob")
	if err != nil {
		fmt.Println(err)
		return activities, err
	}
	defer dataFile.Close()
	dataDecoder := gob.NewDecoder(dataFile)
	err = dataDecoder.Decode(&activities)

	if err != nil {
		fmt.Println(err)
		return activities, err
	}
	return activities, nil
}

func init() {
	rootCmd.AddCommand(indexCmd)
	rootCmd.AddCommand(tickerCmd)
}
