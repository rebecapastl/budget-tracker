package main

import (
	"budget-tracker/export"
	"budget-tracker/process"
	"budget-tracker/standardize"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	fmt.Println("Budget tracker")

	argsWithoutProg := os.Args[1:]
	fmt.Println("argsWithoutProg")
	fmt.Println(argsWithoutProg)

	month := time.Now().AddDate(0, -1, 0).Format("2006-01")
	if len(argsWithoutProg) > 0 {
		month = argsWithoutProg[0]
	}

	
	fmt.Println(month)

	// convert card records
	creditCardExpenses, err := standardize.ExtractCreditCardRecords(month)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(creditCardExpenses)

	// convert account records
	accountExpenses, err := standardize.ExtractAccountRecords(month)
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println(accountExpenses)

	expenseRecords, classifyManually := process.ClassifyInputRecord(creditCardExpenses, accountExpenses)
	fmt.Println("expenseRecords")
	fmt.Println(expenseRecords)
	fmt.Println("classifyManually")
	fmt.Println(classifyManually)

	process.InputMissingInfo(classifyManually)

	fmt.Println("classifyManually")
	fmt.Println(classifyManually)

	// export into CSV
	err = export.ExportToCSV(month, expenseRecords, classifyManually)
	if err != nil {
		log.Fatalf("failed to export to CSV: %s", err)
	}

	// export into Excel
	err = export.ExportToExcel(month, expenseRecords, classifyManually)
	if err != nil {
		log.Fatalf("failed to export to Excel: %s", err)
	}

}
