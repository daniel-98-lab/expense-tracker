package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var monthSummary int
var categorySummary string

// summaryExpense is a Cobra command to display a summary of all expenses, optionally filtered by specific criteria or month.
var summaryExpense = &cobra.Command{
	Use:   "summary",
	Short: "View summary expense with filters",
	Run: func(cmd *cobra.Command, args []string) {

		summary, err := expenseTracker.GetSummaryExpense(monthSummary, categorySummary)

		if err != nil {
			fmt.Printf("Error Summary expense: %s\n", err)
		}

		fmt.Printf("Total Summary: %.2f\n", summary)
	},
}

func init() {

	summaryExpense.Flags().IntVarP(&monthSummary, "month", "m", -1, "Month to filter")
	summaryExpense.Flags().StringVarP(&categorySummary, "category", "c", "", "Category to filter")

	rootCmd.AddCommand(summaryExpense)
}
