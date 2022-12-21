package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	traceCmd = &cobra.Command {
		Use: "trace",
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
		Run: func(cmd *cobra.Command, args []string){
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


type Ip struct {
	IP       string `json:"ip"`
    HOSTNAME string `json:"hostname"`
	CITY     string `json:"city"`
	REGION   string `json:"region"`
	COUNTRY  string `json:"country"`
	LOCATION string `json:"loc"`
    ORG      string `json:"org"`
	POSTAL   string `json:"postal"`
    TIMEZONE string `json:"timezone"`
}

func getData(url string) []byte {

	response, err := http.Get(url)
	if err != nil {
		log.Println("Unable to get the response")
	}

	responseByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Unable to read the response")
	}

	return responseByte
}


func showData(ip string) {

	url := "http://ipinfo.io/" + ip + "/geo"
	responseByte := getData(url)

	data := Ip{}

	err := json.Unmarshal(responseByte, &data)
	if err != nil {
		log.Println("Unable to unmarshal the response")
	}

	c := color.New(color.FgRed).Add(color.Underline).Add(color.Bold)
	c.Println("IP DATA FOUND :")

	display(data)

}


func display(data Ip) {

	fmt.Printf("IP: 		%s\n", data.IP)
	fmt.Printf("HOSTNAME:       %s\n", data.HOSTNAME)
	fmt.Printf("CITY: 		%s\n", data.CITY)
	fmt.Printf("REGION: 	%s\n", data.REGION)
	fmt.Printf("COUNTRY: 	%s\n", data.COUNTRY)
	fmt.Printf("COORDINATES: 	%s\n", data.LOCATION)
	fmt.Printf("ORGANIZAION: 	%s\n", data.ORG)
	fmt.Printf("TIMEZONE: 	%s\n", data.TIMEZONE)
	fmt.Printf("POSTAL CODE: 	%s\n", data.POSTAL)

}