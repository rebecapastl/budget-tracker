package process

import (
	"budget-tracker/models"
	"budget-tracker/utils"
	"fmt"
	"strings"
)

func ClassifyInputRecord(creditCardExpenses, accountExpenses []models.BankRecord) ([]models.ExpenseRecord, []models.ExpenseRecord) {
	var expenseRecords []models.ExpenseRecord
	var classifyManually []models.ExpenseRecord

	for _, record := range accountExpenses {
		purpose := findPurpose(record)
		category := findCategory(purpose)

		expenseRecord := models.ExpenseRecord{
			Date:     record.Date,
			Value:    record.Value,
			Category: category,
			Purpose:  purpose,
			Company:  utils.SanitizeCompanyName(record.Company),
		}

		fmt.Println(expenseRecord)

		if expenseRecord.Category == "" || expenseRecord.Purpose == "" {
			classifyManually = append(classifyManually, expenseRecord)
		} else {
			expenseRecords = append(expenseRecords, expenseRecord)
		}
	}

	//creditCardExpenses has only the company column
	for _, record := range creditCardExpenses {
		purpose := findPurpose(record)
		category := findCategory(purpose)

		expenseRecord := models.ExpenseRecord{
			Date: record.Date,
			Value: record.Value,
			Category: category,
			Purpose: purpose,
			Company: utils.SanitizeCompanyName(record.Company),
		}

		if expenseRecord.Category == "" || expenseRecord.Purpose == "" {
			classifyManually = append(classifyManually, expenseRecord)
		} else {
			expenseRecords = append(expenseRecords, expenseRecord)
		}
	}

	return expenseRecords, classifyManually

}

// purpose to category string
var categoryMap = map[string]string{
	"Groceries/Supplies":  "Home",
	"Meal":                "Home",
	"Rent":                "Home",
	"Heating":             "Home",
	"Mobile":              "Home",
	"Power":               "Home",
	"Internet":            "Home",
	"Maintenance/Repairs": "Home",

	"Food":         "Pets",
	"Pet medicine": "Pets",
	"Pet shop":     "Pets",
	"Grooming":     "Pets",
	"Veterinary":   "Pets",

	"Movies":              "Entretainment",
	"Computer":            "Entretainment",
	"Drawing":             "Entretainment",
	"Sports/events/shows": "Entretainment",
	"Games":               "Entretainment",
	"Books":               "Entretainment",
	"Comics":              "Entretainment",
	"Travel":              "Entretainment",
	"Stream":              "Entretainment",

	"Enrollment":     "Study",
	"Course":         "Study",
	"Study books":    "Study",
	"Study supplies": "Study",

	"Other":      "Extra",
	"Rounding":   "Extra",
	"Donation":   "Extra",
	"Gifts":      "Extra",
	"Withdrawal": "Extra",

	"Tax Return": "Taxes",
	"Fee":        "Taxes",

	"WealthSimple": "Investment",
	"Savings":      "Investment",

	"Clothing": "Personal",
	"Hair":     "Personal",
	"Beauty":   "Personal",
	"Hygene":   "Personal",
	"Laundry":  "Personal",

	"MSP":                "Health",
	"Medicine":           "Health",
	"First aid/supplies": "Health",

	"Transit":         "Commuting",
	"Gas":             "Commuting",
	"Parking":         "Commuting",
	"Car maintenance": "Commuting",
	"Ticket":          "Commuting",
	"Insurance":       "Commuting",
	"Car rent":        "Commuting",
	"Ride":            "Commuting",

	"Salary":     "Income",
	"Investment": "Income",
	"EI Benefit": "Income",
	"Tax Refund": "Income",

	"Transfer to Credit Card": "Accounts Maintenance",
}

// granular
var purposeMap = map[string]string{
	// "Groceries/Supplies",
	"Canadian Superstore": "Groceries/Supplies",
	"Buy Low Foods":       "Groceries/Supplies",
	"Tru Earth":           "Groceries/Supplies",
	// "Meal",
	"Skip The Dishes": "Meal",
	// "Rent",
	// "Heating",
	// "Mobile",
	"Rogers": "Mobile",
	// "Power",
	// "Internet",
	// "Maintenance/Repairs",

	// "Food",
	// "Pet medicine",
	// "Pet shop",
	// "Grooming",
	// "Veterinary",

	// "Movies",
	// "Computer",
	// "Drawing",
	// "Sports/events/shows",
	// "Games",
	// "Books",
	// "Comics",
	// "Travel",
	// "Stream",
	"Crave": "Stream",

	// "Enrollment",
	// "Course",
	// "Study books",
	// "Study supplies",

	// "Other",
	"POS Purchase OPOS Amazon.ca Prime Membamazon.ca/pBCCA": "Other",
	// "Rounding",
	// "Donation",
	// "Gifts",
	// "Withdrawal",
	"WITHDRAWAL": "Withdrawal",

	// "Tax Return",
	// "Fee",

	// "WealthSimple",
	"Investment Wealthsimple Investments Inc.": "WealthSimple",
	// "Savings",
	"Customer Transfer Dr.": "Savings",

	// "Clothing",
	// "Hair",
	// "Beauty",
	// "Hygene",
	// "Laundry",

	// "MSP",
	"Revenue Services BC": "MSP",
	// "Medicine",
	// "First aid/supplies",

	// "Transit",
	"Compass Card": "Transit",
	// "Gas",
	// "Parking",
	// "Car maintenance",
	// "Ticket",
	// "Insurance",
	// "Car rent",
	// "Ride",
	"Uber": "Ride",
	"Lyft": "Ride",

	//"Salary",
	"Payroll Deposit": "Salary",
	//"Investment",
	//"Tax Refund",
	"Fed-Prov/Terr CANADA": "Tax Refund",

	//"Credit Card",
	"Credit Card": "Transfer to Credit Card",
}

func findPurpose(record models.BankRecord) string {
	purpose := purposeMap[strings.TrimSpace(record.Purpose+" "+record.Company)]
	return purpose
}

func findCategory(purpose string) string {
	category := categoryMap[purpose]
	return category
}
