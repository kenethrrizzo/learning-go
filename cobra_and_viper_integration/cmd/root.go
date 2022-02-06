/*
Copyright © 2022 KENETH RIERA KENETHRR@OUTLOOK.COM
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var name string

var rootCmd = &cobra.Command{
	Use:   "cobra_and_viper_integration",
	Short: "A brief description of your application",
	Long: `Cobra is a CLI library for Go that empowers applications.
		This application is a tool to generate the needed files
		to quickly create a Cobra application.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "add name")
	viper.BindPFlag("name", rootCmd.PersistentFlags().Lookup("name"))
}


