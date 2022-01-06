package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/spf13/viber"
)

var cfgFile string


var rootCmd = &cobra.Command{
	Use: "ip-tool"
	Short: "IP Tool CLI in GO"
	Long: "IP Tool CLI application writtenin GO"
}


func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file default")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toogle")

}



func initConfig(){
	if cfgFile != ""{
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".ip-tool")

	}
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

}
