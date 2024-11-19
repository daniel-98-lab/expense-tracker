package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var typeFile string

var exportExpense = &cobra.Command{
	Use:   "export",
	Short: "Export expenses to file",
	Run: func(cmd *cobra.Command, args []string) {

		err := expenseTracker.ExportExpense(typeFile)

		if err != nil {
			fmt.Printf("Error Export expense: %s\n", err)
		}

		fmt.Printf("Export completed\n")
	},
}

func init() {

	exportExpense.Flags().StringVarP(&typeFile, "type", "t", "json", "Type file to export(json or csv)")

	rootCmd.AddCommand(exportExpense)
}
