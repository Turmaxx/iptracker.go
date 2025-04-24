package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	traceCmd = &cobra.Command{
		Use:   "trace",
		Short: "Trace & Find infomation about an IP address",
		Long: ` Find Information about the IP and formats it as such

			IP 		
            HOSTNAME
			CITY 		
			REGION 	
			COUNTRY 	
			COORDINATES 
            ORGANIZAION
			TIMEZONE 	
			POSTAL CODE 	

		`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				for _, ip := range args {
					showData(ip)
				}
			} else {
				fmt.Println("Please provide the IP address")
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(traceCmd)
}

type InternetInfo struct {
	IP           string `json:"ip"`
	CITY         string `json:"city"`
	REGION       string `json:"region"`
	COUNTRY      string `json:"country"`
	LOCATION     string `json:"loc"`
	ORGANIZATION string `json:"org"`
	TIMEZONE     string `json:"timezone"`
}

func getData(url string) []byte {

	response, err := http.Get(url)
	if err != nil {
		log.Println("Unable to get the response")
	}

	responseByte, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Unable to read the response")
	}

	return responseByte
}

func showData(ip string) {

	url := "http://ipinfo.io/" + ip + "/geo"
	responseByte := getData(url)

	data := InternetInfo{}

	err := json.Unmarshal(responseByte, &data)
	if err != nil {
		log.Println("Unable to unmarshal the response")
	}

	display(data)
}

func display(data InternetInfo) {
	red := color.New(color.FgRed).Add(color.Underline).Add(color.Bold)
	red.Println("IP DATA FOUND:")
	green := color.New(color.FgHiGreen)
	green.Printf("IP: %s\n", data.IP)
	green.Printf("CITY: %s\n", data.CITY)
	green.Printf("REGION: %s\n", data.REGION)
	green.Printf("COUNTRY: %s\n", data.COUNTRY)
	green.Printf("LOCATION: %s\n", data.LOCATION)
	green.Printf("ORGANIZATION: %s\n", data.ORGANIZATION)
	green.Printf("POSTAL CODE: %s\n", data.TIMEZONE)
}
