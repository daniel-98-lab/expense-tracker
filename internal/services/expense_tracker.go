package services

import (
	"encoding/json"
	"expense-tracker/internal/models"
	"fmt"
	"os"
	"time"
)

// ExpenseTracker provides methods to manage expenses
type ExpenseTracker struct {
	filePath string
}

// Constructor initializes and returns a new ExpenseTracker with the provided file path.
func Constructor(filepath string) *ExpenseTracker {
	return &ExpenseTracker{filePath: filepath}
}

func (s *ExpenseTracker) CreateExpense(description string, amount float64, date string) (string, error) {
	expenses, err := s.LoadExpenses()
	if err != nil {
		return "An error has occurred", err
	}

	//TODO: VERIFY DE DATE IS IN THE CORRECT FORMAT, AND IS A CORRECT DATE

	id := len(expenses) + 1

	newExpense := models.Expense{
		ID:          id,
		Amount:      amount,
		Date:        date,
		Description: description,
	}

	expenses = append(expenses, newExpense)
	err = s.SaveExpenses(expenses)
	return fmt.Sprintf("Expense added successfully: %d", id), err
}

func (s *ExpenseTracker) DeleteExpense(id int) error {
	expenses, err := s.LoadExpenses()
	if err != nil {
		return err
	}

	for i, expense := range expenses {
		if expense.ID == id {
			expenses = append(expenses[:i], expenses[:i+1]...)
			return s.SaveExpenses(expenses)
		}
	}
	return nil
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

func (s *ExpenseTracker) UpdateExpense(id int, description string, amount float64, date string) (string, error) {
	expenses, err := s.LoadExpenses()
	if err != nil {
		return fmt.Sprintln("error to load expenses"), err
	}

	for i, expense := range expenses {
		if expense.ID == id {
			expenses[i].Amount = amount
			expenses[i].Date = date
			expenses[i].Description = description
			err = s.SaveExpenses(expenses)
			return fmt.Sprintf("error to update expense: %d", id), err
		}
	}
	return fmt.Sprintf("Expense updated successfully ID: %d", id), nil
}

func (s *ExpenseTracker) SaveExpenses(expenses []models.Expense) error {
	file, err := os.Create(s.filePath)

	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(expenses)
}

func (s *ExpenseTracker) ShowSummary(month string) (string, error) {
	expenses, err := s.LoadExpenses()
	if err != nil {
		return fmt.Sprintln("error to load expenses"), err
	}

	total := 0.0

	for _, expense := range expenses {
		if month != "" {
			expenseDate, err := time.Parse("2006-01-02", expense.Date)
			if err != nil {
				return fmt.Sprintln("error parsing expense date"), err
			}
			if fmt.Sprintf("%04d-%02d", expenseDate.Year(), int(expenseDate.Month())) != month {
				continue
			}
		}

		total += expense.Amount
	}

	return fmt.Sprintf("Total expenses: %.2f", total), nil
}
