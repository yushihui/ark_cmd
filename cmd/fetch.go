package cmd

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
)

const arkPrefix = "https://ark-funds.com/wp-content/fundsiteliterature/csv/"
const folderPrefix = "c://ark/history/"

var arkFund = map[string]string{
	"ARKK": "ARK_INNOVATION_ETF_ARKK_HOLDINGS.csv",
	"ARKQ": "ARK_AUTONOMOUS_TECHNOLOGY_&_ROBOTICS_ETF_ARKQ_HOLDINGS.csv",
	"ARKW": "ARK_NEXT_GENERATION_INTERNET_ETF_ARKW_HOLDINGS.csv",
	"ARKG": "ARK_GENOMIC_REVOLUTION_MULTISECTOR_ETF_ARKG_HOLDINGS.csv",
	"ARKF": "ARK_FINTECH_INNOVATION_ETF_ARKF_HOLDINGS.csv",
	"PRNT": "THE_3D_PRINTING_ETF_PRNT_HOLDINGS.csv",
	"IZRL": "ARK_ISRAEL_INNOVATIVE_TECHNOLOGY_ETF_IZRL_HOLDINGS.csv",
}

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "fetch ARK fund profilio",
	RunE: func(cmd *cobra.Command, args []string) error {
		doDownload()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
}

func doDownload() {
	currentTime := time.Now()
	timeStr := currentTime.Format("2006-01-02")
	for k, v := range arkFund {
		err := downloadFile("data/"+k+"/"+timeStr+".csv", arkPrefix+v)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(timeStr, " successfully downloaded ", k)
		}
		time.Sleep(2 * time.Second)
	}
}

func downloadFile(filepath string, url string) error {

	if _, err := os.Stat(filepath); err == nil {
		return errors.New("file is already exist " + filepath)
	}
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}
