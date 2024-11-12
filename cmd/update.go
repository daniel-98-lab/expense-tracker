package main

import (
	"fmt"

	"github.com/daniel-98-lab/expense-tracker/internal/models"
	"github.com/spf13/cobra"
)

var updateData models.Expense

var updateExpense = &cobra.Command{
	Use:   "update",
	Short: "Update an expense",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Updating expense: %v\n", updateData.ID)
	},
}

func init() {
	updateExpense.Flags().IntVarP(&updateData.ID, "id", "i", 0, "Identification of the expense")
	updateExpense.Flags().StringVarP(&updateData.Description, "description", "d", "", "Description of the expense")
	updateExpense.Flags().Float64VarP(&updateData.Amount, "amount", "a", 0.0, "Amount of the expense")
	updateExpense.Flags().StringVarP(&updateData.Date, "date", "D", "", "Date of the expense")
	updateExpense.Flags().StringVarP(&updateData.Category, "category", "c", "", "Category of the expense")
	updateExpense.MarkFlagRequired("id")

	rootCmd.AddCommand(updateExpense)
}