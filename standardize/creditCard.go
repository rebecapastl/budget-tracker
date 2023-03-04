package standardize

import (
	"budget-tracker/models"
	"budget-tracker/utils"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func ExtractCreditCardRecords(month string) ([]models.BankRecord, error) {
	// process card records
	// open card file
	cardFile, err := os.Open("files/" + month + "_credit-card.csv")
	if err != nil {
		fmt.Println(err)
	}

	// create reader
	cardReader := csv.NewReader(cardFile)

	// read card records
	cardRecords, err := cardReader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	var creditCardExpenses []models.BankRecord
	for _, record := range cardRecords {

		date := utils.FormatDate(record[0])

		value, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		company := utils.SanitizeCompanyName(record[1])

		expense := models.BankRecord{
			Date:    date,
			Value:   value,
			Purpose: "",
			Company: company,
		}
		creditCardExpenses = append(creditCardExpenses, expense)

	}

	return creditCardExpenses, nil

}

