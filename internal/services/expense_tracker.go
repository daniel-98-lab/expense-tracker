package services

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/daniel-98-lab/expense-tracker/internal/models"
)

type ExpenseTracker struct {
	filePath string
}

// constructor that initializes the ExpenseTracker service, its sets the file path used to save the expense data
func NewExpenseService(filepath string) *ExpenseTracker {
	return &ExpenseTracker{filePath: filepath}
}

// function to create a new expense
// First verifies the date is correct, loads existing expenses from the root file,
// Creates a new expense object, and finally call SaveExpenses to save it to the file
func (s *ExpenseTracker) CreateExpense(description string, amount float64, date string, category string) (models.Expense, error) {
	if _, err := parseDate(date); err != nil {
		return models.Expense{}, err
	}

	expenses, err := s.LoadExpenses()
	if err != nil {
		return models.Expense{}, err
	}
	id := len(expenses) + 1

	newExpense := models.Expense{
		ID:          id,
		Amount:      amount,
		Date:        date,
		Description: description,
		Category:    category,
	}

	expenses = append(expenses, newExpense)
	result := s.SaveExpenses(expenses)
	return newExpense, result
}

// function to remove an expense
// Loads existing expenses,  checks if the expense exists, and if it doesn't, displays an error message.
// If the expense exists, it removes the expense and saves the updated expense data.
func (s *ExpenseTracker) DeleteExpense(id int) error {
	expenses, err := s.LoadExpenses()

	if err != nil {
		return err
	}

	if !Any(expenses, func(exp models.Expense) bool { return exp.ID == id }) {
		err := fmt.Errorf("there is no charge with this ID: %d", id)
		return err
	}

	for i := range expenses {
		if expenses[i].ID == id {
			expenses = append(expenses[:i], expenses[i+1:]...)
			return s.SaveExpenses(expenses)
		}
	}

	return nil
}

// Exports data to a file. Only two formats are supported: JSON or CSV.
func (s *ExpenseTracker) ExportExpense(typeExport string) error {

	switch typeExport {
	case "json":
		return CopyFile(s.filePath)
	case "csv":
		return JSONToCSV(s.filePath)
	default:
		return fmt.Errorf("typeExport incorrect, write json or csv")
	}
}

// / function to get the summary of all expenses,
func (s *ExpenseTracker) GetSummaryExpense(month int, category string) (float64, error) {
	var expenses []models.Expense
	var err error
	var summary float64
	expenses, err = s.LoadExpensesByCategory(category)

	if err != nil {
		return summary, err
	}

	for _, expense := range expenses {
		if month == -1 || isMatchingMonth(expense.Date, month) {
			summary += expense.Amount
		}
	}

	return summary, nil
}

// Retrieves the summary of all expenses, optionally filtered by month and category.
// If the month is -1, all expenses are included regardless of the month.
func (s *ExpenseTracker) LoadExpenses() ([]models.Expense, error) {

	file, err := os.Open(s.filePath)
	if err != nil {

		if os.IsNotExist(err) {

			emptyExpenses := []models.Expense{}
			if err := s.SaveExpenses(emptyExpenses); err != nil {
				return nil, fmt.Errorf("error creating file: %v", err)
			}

			return emptyExpenses, nil
		}
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	expenses := []models.Expense{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&expenses)
	if err != nil {
		return nil, fmt.Errorf("error decoding expenses.json: %v", err)
	}

	return expenses, nil
}

// Loads all expenses and filters them by the specified category.
// If no category is provided, all expenses are returned.
func (s *ExpenseTracker) LoadExpensesByCategory(category string) ([]models.Expense, error) {
	expenses, err := s.LoadExpenses()

	if err != nil {
		return nil, err
	}

	if category == "" {
		return expenses, nil
	}

	// Filter expenses by the provided category
	var filteredExpenses []models.Expense
	for _, expense := range expenses {
		if expense.Category == category {
			filteredExpenses = append(filteredExpenses, expense)
		}
	}
	return filteredExpenses, nil
}

// Updates an existing expense with the given ID. It checks if the expense exists,
// validates the provided date, and then updates the expense details (amount,
// category, date, description). Finally, it saves the updated expenses to the file.
func (s *ExpenseTracker) UpdateExpense(id int, description string, amount float64, date string, category string) error {
	if _, err := parseDate(date); err != nil {
		return err
	}

	expenses, err := s.LoadExpenses()
	if err != nil {
		return err
	}

	if !Any(expenses, func(exp models.Expense) bool { return exp.ID == id }) {
		err := fmt.Errorf("there is no charge with this ID: %d", id)
		return err
	}

	for i := range expenses {
		if expenses[i].ID == id {
			expenses[i].Amount = amount
			expenses[i].Category = category
			expenses[i].Date = date
			expenses[i].Description = description
			err = s.SaveExpenses(expenses)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Saves the updated list of expenses to the file specified by the file path.
func (s *ExpenseTracker) SaveExpenses(expenses []models.Expense) error {
	file, err := os.Create(s.filePath)

	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(expenses)
}

//GLOBAL FUNCTIONS

// Return true if expense exist in the slice.
func Any(expenses []models.Expense, condition func(models.Expense) bool) bool {
	for _, expense := range expenses {
		if condition(expense) {
			return true
		}
	}
	return false
}

// function to copy json files
func CopyFile(filePath string) error {

	//open origin file
	sourceFile, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer sourceFile.Close()

	//create destination File
	destinationFile, err := os.Create("exports/expenses.json")
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer destinationFile.Close()

	//Copy content file
	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return fmt.Errorf("failed to copy file content: %w", err)
	}

	// Forzar the content save of the disk
	err = destinationFile.Sync()
	if err != nil {
		return fmt.Errorf("failed to sync destination file: %w", err)
	}

	return nil
}

// Creates a new CSV file and exports the expenses from the original file into the new CSV file.
func JSONToCSV(filePath string) error {
	fileOrigin, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open JSON file: %w", err)
	}
	defer fileOrigin.Close()

	var expenses []models.Expense
	if err := json.NewDecoder(fileOrigin).Decode(&expenses); err != nil {
		return fmt.Errorf("failed to decode JSON: %w", err)
	}

	csvFile, err := os.Create("exports/expenses.csv")
	if err != nil {
		return fmt.Errorf("failed to create CSV file: %w", err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	if err := writer.Write([]string{"ID", "Description", "Amount", "Date", "Category"}); err != nil {
		return fmt.Errorf("failed to write CSV header: %w", err)
	}

	for _, expense := range expenses {
		record := []string{
			fmt.Sprintf("%d", expense.ID),
			expense.Description,
			fmt.Sprintf("%.2f", expense.Amount),
			expense.Date,
			expense.Category,
		}
		if err := writer.Write(record); err != nil {
			return fmt.Errorf("failed to write CSV record: %w", err)
		}
	}

	return nil
}

// Checks if the given date matches the specified month.
func isMatchingMonth(date string, month int) (result bool) {
	parsedDate, err := parseDate(date)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return false
	}
	return int(parsedDate.Month()) == month
}

// Verifies if the provided date is valid.
func parseDate(date string) (time.Time, error) {
	dateParse, err := time.Parse("2006-01-02", date)
	return dateParse, err
}
