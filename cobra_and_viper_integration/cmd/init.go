/*
Copyright © 2022 KENETH RIERA KENETHRR@OUTLOOK.COM
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `Cobra is a CLI library for Go that empowers applications.
		This application is a tool to generate the needed files
		to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
		fmt.Println(viper.GetInt("port"))
		if name != "" {
			fmt.Println("Name:", viper.GetString("name"))
		}
	},
}

func init() {
	initCmd.Flags().IntP("port", "p", 3000, "add server's port")
	viper.BindPFlag("port", initCmd.Flags().Lookup("port"))

	rootCmd.AddCommand(initCmd)
}
