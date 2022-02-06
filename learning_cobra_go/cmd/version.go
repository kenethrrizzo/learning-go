/*
Copyright © 2022 Keneth Riera kenethrr@outlook.com
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var task string

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "A brief description of your command",
	Long: `Cobra is a CLI library for Go that empowers applications.
			This application is a tool to generate the needed files
			to quickly create a Cobra application.`,
	// Before
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("Before")
	},
	Run: func(cmd *cobra.Command, args []string) {
		path, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		fmt.Println(path)
		if verbose {
			fmt.Println("More logs...")
		}
		if name != "" {
			fmt.Println("Hello,", name)
		}
		fmt.Println("Task:", task)
	},
	// After
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("After")
	},
	Args: cobra.MaximumNArgs(2),
}

func init() {
	versionCmd.Flags().StringVarP(&task, "task", "t", "", "Add new task.")

	//versionCmd.MarkFlagRequired("task")

	rootCmd.AddCommand(versionCmd)
}
