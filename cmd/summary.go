package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var summaryExpense = &cobra.Command{
	Use:   "summary",
	Short: "View summary expense with filters",
	Run: func(cmd *cobra.Command, args []string) {
		summary := 12
		fmt.Printf("Total expenses: $%v\n", summary)
	},
}

func init() {
	var month int
	var category string

	summaryExpense.Flags().IntVarP(&month, "month", "m", 0, "Month to filter")
	summaryExpense.Flags().StringVarP(&category, "category", "c", "", "Category to filter")

	rootCmd.AddCommand(summaryExpense)
}
