package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Start Prometheus metrics endpoint",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting metrics endpoint on :9000/metrics...")
		// Start your Prometheus monitoring logic here
	},
}

func init() {
	rootCmd.AddCommand(monitorCmd)
}
