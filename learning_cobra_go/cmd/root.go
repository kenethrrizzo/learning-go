/*
Copyright © 2022 Keneth Riera kenethrr@outlook.com
*/
package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var verbose bool
var name string

var rootCmd = &cobra.Command{
	Use:   "appcli",
	Short: "A brief description of your application",
	Long: `Cobra is a CLI library for Go that empowers applications.
			This application is a tool to generate the needed files
			to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, world")
		if verbose {
			fmt.Println("I love the world")
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Show more info.")
	rootCmd.PersistentFlags().StringVarP(&name, "name", "n", "user", "Set username.")
	rootCmd.MarkPersistentFlagRequired("name")
}
