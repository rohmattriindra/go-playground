package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

var getIP = &cobra.Command{
	Use:   "ip",
	Short: "This command to gathering IP Public",
	Long:  "This get command will call GitHub",
	Run: func(cmd *cobra.Command, args []string) {
		//  var ipName = "ip"

		//  if len(args) >1 && args[0] != ""{
		// 	 ipName = args[0]
		//  }

		URL := "https://api.ipify.org"

		fmt.Println("Tryng to heck your IP Address")

		response, err := http.Get(URL)

		fmt.Println(response)

		if err != nil {
			fmt.Println(err)
			// handle error
		}

		defer response.Body.Close()

		if response.StatusCode == 200 {
			fmt.Println("Sucessfull")
		} else {
			fmt.Println("Connection Failed")
		}
	},
}

func init() {
	rootCmd.AddCommand(getIP)
}
