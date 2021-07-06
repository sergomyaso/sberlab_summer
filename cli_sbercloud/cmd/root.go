package cmd

import (
"fmt"
"github.com/spf13/cobra"
"os"
)

var developer string

var rootCmd = &cobra.Command{
	Use:   "broo",
	Short: "cli for sbercloud",
	Long:  `special cli for work with sbercloud API `,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&developer, "developer", "sergomyaso", "Developer name.")
}

func initConfig() {
	developer, _ := rootCmd.Flags().GetString("developer")
	if developer != "" {
		fmt.Println("Developer:", developer)
	}
}

