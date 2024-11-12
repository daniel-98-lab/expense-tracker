package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listExpense = &cobra.Command{
	Use:   "list",
	Short: "List all expenses",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Expenses: \n")
	},
}

func init() {
	rootCmd.AddCommand(listExpense)
}
