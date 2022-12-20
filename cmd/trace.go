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
			CITY 		
			REGION 	
			COUNTRY 	
			COORDINATES 
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
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Location string `json:"loc"`
	Timezone string `json:"timezone"`
	Postal   string `json:"postal"`
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
	fmt.Printf("CITY: 		%s\n", data.City)
	fmt.Printf("REGION: 	%s\n", data.Region)
	fmt.Printf("COUNTRY: 	%s\n", data.Country)
	fmt.Printf("COORDINATES: 	%s\n", data.Location)
	fmt.Printf("TIMEZONE: 	%s\n", data.Timezone)
	fmt.Printf("POSTAL CODE: 	%s\n", data.Postal)

}