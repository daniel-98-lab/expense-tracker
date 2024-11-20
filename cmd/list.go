package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var categoryList string

// listExpense is a Cobra Command to display all expenses,optionally filtered by specific criteria
var listExpense = &cobra.Command{
	Use:   "list",
	Short: "List all expenses",
	Run: func(cmd *cobra.Command, args []string) {
		expenses, err := expenseTracker.LoadExpensesByCategory(categoryList)

		if err != nil {
			fmt.Printf("Error to load expenses %s\n", err)
		}

		if len(expenses) == 0 {
			fmt.Println("Wow, it looks like you don't have any expense")
		}

		for _, expense := range expenses {
			fmt.Printf("******************************************\n")
			fmt.Printf("*******Expense id: %v \n", expense.ID)
			fmt.Printf("*******Expense Description: %v \n", expense.Description)
			fmt.Printf("*******Expense Amount: %v \n", expense.Amount)
			fmt.Printf("*******Expense Date: %v \n", expense.Date)
			fmt.Printf("*******Expense Category: %v \n", expense.Category)
			fmt.Printf("******************************************\n")
		}
	},
}

func init() {
	listExpense.Flags().StringVarP(&categoryList, "category", "c", "", "Category to filter")

	rootCmd.AddCommand(listExpense)
}
