package models

type BankRecord struct {
	Date    string
	Value   float64
	Purpose string
	Company string
}

type ExpenseRecord struct {
	Date    string
	Value   float64
	Category	string
	Purpose string
	Company string
}

