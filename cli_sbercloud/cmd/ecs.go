package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ecsCommand = &cobra.Command{
	Use:   "ecs",
	Short: "interaction with Elastic Cloud Server  ",
	Long: `Managing and getting statistics on Elastic Cloud Server`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hola, brooo :)")
	},
}

var StatisticSubcommand = &cobra.Command{
	Use:   "statistic",
	Short: "get statistic about servers",
	Long: `Managing and getting statistics on Elastic Cloud Server`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hola, brooo :)")
	},
}

func init() {
	rootCmd.AddCommand(ecsCommand)
}