package ark

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strconv"
)

func fileLoc(fund, filename string) string {
	return "data/" + fund + "/" + filename + ".csv"
}

func deltaChange(fund, from, to string) map[string]*Security {
	securities := csvReadByLine(fileLoc(fund, from))
	recordFile, err := os.Open(fileLoc(fund, to))
	if err != nil {
		fmt.Println("An error encountered ::", err)
	}
	r := csv.NewReader(recordFile)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}
		if len(record[0]) == 0 || record[0][0] < '0' || record[0][0] > '9' {
			continue
		}
		shares, _ := strconv.ParseFloat(record[5], 64)
		value, _ := strconv.ParseFloat(record[6], 64)
		w, _ := strconv.ParseFloat(record[7], 64)
		price := value / shares
		if s, ok := securities[record[4]]; ok {
			s.Delta = shares - s.Shares
			s.Price = price
			s.Weight = w
		} else {
			s := Security{
				Fund:       record[1],
				Name:       record[2],
				TickerCode: record[4],
				Shares:     shares,
				Delta:      shares,
				Price:      price,
				IsNew:      true,
				Weight:     w,
			}
			securities[record[4]] = &s
		}
	}
	for k, v := range securities {
		if v.Delta == 0 {
			delete(securities, k)
		}

	}
	return securities
}

// FundActivity calc fund delta changes
func FundActivity(fund string, from string, to string) {
	securities := deltaChange(fund, from, to)
	prettyPrint(securities)
}

func prettyPrint(securityMap map[string]*Security) {
	fmt.Printf("%-10s %-30s %20s %20s %20s\n", "Direction", "Name", "Shares", "Value", "Weight(%)")
	securities := make([]*Security, 0, len(securityMap))
	for _, v := range securityMap {
		securities = append(securities, v)
	}
	sort.Sort(ByValue(securities))
	var direction string
	for _, s := range securities {
		direction = "buy"
		if s.IsNew {
			direction = "*buy"
		} else if s.Delta < 0 {
			direction = "sell"
		}
		fmt.Printf("%-10s %-30s %20s %20s %20.2f\n", direction, s.Name, prettyNumber(s.Delta), prettyNumber(s.Delta*s.Price), s.Weight)
	}
}

func prettyNumber(number float64) string {
	n := math.Abs(number)
	if n < 1000 {
		return fmt.Sprintf("%.0f", number)
	}
	exp := (int64)(math.Log(n) / math.Log(1000))
	result := fmt.Sprintf("%.2f %c", n/math.Pow(1000, (float64)(exp)), "KMB"[exp-1])
	if number < 0 {
		result = "-" + result
	}
	return result
}

func csvReadByLine(file string) map[string]*Security {
	csvfile, err := os.Open(file)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	securities := make(map[string]*Security)
	r := csv.NewReader(csvfile)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil || len(record[0]) == 0 || record[0][0] < '0' || record[0][0] > '9' {
			continue
		}
		shares, _ := strconv.ParseFloat(record[5], 64)
		s := Security{
			Fund:       record[1],
			Name:       record[2],
			TickerCode: record[4],
			Shares:     shares,
			Delta:      0,
		}

		securities[record[4]] = &s
	}
	return securities
}

// Profilio fund profilio
func Profilio(fund string) {
	fileName := getLatestFile("data/" + fund)
	csvfile, err := os.Open(fileName)
	if err != nil {
		log.Fatalln("Couldn't open the data file", err)
	}
	r := csv.NewReader(csvfile)
	fmt.Printf("%-10s %-28s %20s %20s\n", "Date", "Name", "Ticker", "Weight(%)")
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil || len(record[0]) == 0 || record[0][0] < '0' || record[0][0] > '9' {
			continue
		}
		fmt.Printf("%-10s %-28s %20s %20s\n", record[0], record[2], record[3], record[7])
	}
}

func getLatestFile(folder string) string {
	file := ""
	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if path > file {
			file = path
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return file
}

// AllFundsActivity  ark transaction activity
func AllFundsActivity(from, to string) {
	securities := make(map[string]*Security)
	for k := range ArkFunds {
		fundSecurities := deltaChange(k, from, to)
		mergeTo(securities, fundSecurities)
	}
	prettyPrint(securities)
}

func mergeTo(target, from map[string]*Security) {
	for k, v := range from {
		if s, ok := target[k]; ok {
			s.Delta += v.Delta
		} else {
			target[k] = v
		}
	}
}
