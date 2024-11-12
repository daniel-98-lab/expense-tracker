package main

import (
	"fmt"

	"github.com/daniel-98-lab/expense-tracker/internal/models"
	"github.com/spf13/cobra"
)

var addData models.Expense

var addExpense = &cobra.Command{
	Use:   "add",
	Short: "Add a new expense",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Adding expense: %s - %.2f\n", addData.Description, addData.Amount)
	},
}

func init() {
	addExpense.Flags().StringVarP(&addData.Description, "description", "d", "", "Description of the expense")
	addExpense.Flags().Float64VarP(&addData.Amount, "amount", "a", 0.0, "Amount of the expense")
	addExpense.Flags().StringVarP(&addData.Date, "date", "D", "", "Date of the expense")
	addExpense.Flags().StringVarP(&addData.Category, "category", "c", "", "Category of the expense")
	addExpense.MarkFlagRequired("description")
	addExpense.MarkFlagRequired("amount")
	addExpense.MarkFlagRequired("date")
	addExpense.MarkFlagRequired("category")

	rootCmd.AddCommand(addExpense)
}
