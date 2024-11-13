package services

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/daniel-98-lab/expense-tracker/internal/models"
)

type ExpenseTracker struct {
	filePath string
}

func NewExpenseService(filepath string) *ExpenseTracker {
	return &ExpenseTracker{filePath: filepath}
}

func (s *ExpenseTracker) CreateExpense(description string, amount float64, date string, category string) (models.Expense, error) {
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

func (s *ExpenseTracker) UpdateExpense(id int, description string, amount float64, date string, category string) error {
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

func isMatchingMonth(date string, month int) (result bool) {
	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return false
	}
	return int(parsedDate.Month()) == month
}
