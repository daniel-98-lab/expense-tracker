package main

import (
	"fmt"
	"os"

	"github.com/daniel-98-lab/expense-tracker/internal/services"
	"github.com/spf13/cobra"
)

var expenseTracker services.ExpenseTracker

// rootCmd is the main command of this app, it serves as the entry point to execute various subcommands
var rootCmd = &cobra.Command{
	Use:   "expense-tracker",
	Short: "Expense Tracker CLI",
	Long:  "An app to track expenses from the command Line",
}

func main() {
	expenseTracker = *services.NewExpenseService("data/expenses.json")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
