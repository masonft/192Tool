package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "192tool",
	Short: "192Tool is a scraper for 192.com",
	Long:  `192Tool is a scraper for 192.com`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Lol")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
