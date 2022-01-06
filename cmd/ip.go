package cmd

 import (
	 "fmt"
	 "io"
	 "net/http"
	 "os"
	 "github.com/spf13/cobra"
 )



 var getIP = &cobra.Command{
	 Use: "ip"
	 Short : "This command to gathering IP Public"
	 Long: "This get command will call GitHub"
	 Run: func(cmd *cobra.Command, args []string){
		//  var ipName = "ip"

		//  if len(args) >1 && args[0] != ""{
		// 	 ipName = args[0]
		//  }

		 URL := "https://ifconfig.co"

		 fmt.Fprintln("Tryng to heck your IP Address")

		 response, err := http.Get(URL)
		 if err != nil {
			 fmt.Println(err)
		 }

		 defer response.Body.Close()

		 if response.StatusCode == 200 {
			 fmt.Println("Sucessfull to connect IFCONFIG.CO")
		 } else {
			 fmt.Println("Connection Failed")
		 }
	}

func init() {
	rootCmd.AddCommand(getIP)
}
