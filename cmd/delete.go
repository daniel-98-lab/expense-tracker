package main

import (
	"fmt"

	"github.com/daniel-98-lab/expense-tracker/internal/models"
	"github.com/spf13/cobra"
)

var deleteData models.Expense

var deleteExpense = &cobra.Command{
	Use:   "delete",
	Short: "Delete an expense",
	Run: func(cmd *cobra.Command, args []string) {
		err := expenseTracker.DeleteExpense(deleteData.ID)

		if err != nil {
			fmt.Printf("Error Delete expense: %s\n", err)
			return
		}

		fmt.Printf("Deleting expense: %v\n", deleteData.ID)
	},
}

func init() {
	deleteExpense.Flags().IntVarP(&deleteData.ID, "id", "i", 0, "Identification of the expense")
	deleteExpense.MarkFlagRequired("id")

	rootCmd.AddCommand(deleteExpense)
}
