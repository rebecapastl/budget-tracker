package standardize

import (
	"budget-tracker/models"
	"budget-tracker/utils"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ExtractAccountRecords(month string) ([]models.BankRecord, error) {
	// process account records
	// open account file
	accountFile, err := os.Open("files/" + month + "_account.csv")
	if err != nil {
		fmt.Println(err)
	}

	// create reader
	accountReader := csv.NewReader(accountFile)

	// read account records
	accountRecords, err := accountReader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	var accountExpenses []models.BankRecord
	for _, record := range accountRecords {

		date := utils.FormatDate(record[0])

		value, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		purpose := strings.TrimSpace(record[3])
		company := strings.TrimSpace(record[4])
		if (purpose == "Customer Transfer Dr." && company == "MB-TRANSFER")|| purpose == "Payroll Deposit" || purpose == "WITHDRAWAL" {
			company = ""
		} else if purpose == "Customer Transfer Dr." && strings.Contains(company, "PC TO"){
			company = ""
			purpose = "Credit Card"
		}

		expense := models.BankRecord{
			Date:    date,
			Value:   value,
			Purpose: purpose,
			Company: company,
		}
		accountExpenses = append(accountExpenses, expense)

	}

	return accountExpenses, nil
}
