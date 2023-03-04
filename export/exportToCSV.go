package export

import (
	"budget-tracker/models"
	"encoding/csv"
	"os"
	"strconv"
)

func ExportToCSV(month string, expenseRecords, classifyManually []models.ExpenseRecord) error {
	csvFile, err := os.Create(month + "_budget.csv")
	if err != nil {
		return err
	}

	csvwriter := csv.NewWriter(csvFile)
	for _, record := range expenseRecords {
		row := []string{record.Date, strconv.FormatFloat(record.Value, 'E', 2, 64), record.Category, record.Purpose, record.Company}
		err = csvwriter.Write(row)
		if err != nil {
			return err
		}
	}

	for _, record := range classifyManually {
		row := []string{record.Date, strconv.FormatFloat(record.Value, 'E', 2, 64), record.Category, record.Purpose, record.Company}
		err = csvwriter.Write(row)
		if err != nil {
			return err
		}
	}

	csvwriter.Flush()
	csvFile.Close()

	return nil
}
