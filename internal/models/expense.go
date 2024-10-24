package models

type Expense struct {
	ID          int     `json:"id"`
	Amount      float64 `json:"amount"`
	Date        string  `json:"date"`
	Description string  `json:"description"`
}
