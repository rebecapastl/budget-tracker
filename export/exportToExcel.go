package export

import (
	"budget-tracker/models"
	"budget-tracker/utils"
	"budget-tracker/process"

	"fmt"
	"strconv"

	"github.com/xuri/excelize"
)

func ExportToExcel(month string, expenseRecords, classifyManually []models.ExpenseRecord) error {
	// Create an Excel file
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// Rename default sheet1
	err := f.SetSheetName("Sheet1", "Despesas")
	if err != nil {
		return err
	}

	// Set title row
	f.SetCellValue("Despesas", "A1", "Date")
	f.SetCellValue("Despesas", "B1", "Company")
	f.SetCellValue("Despesas", "C1", "Category")
	f.SetCellValue("Despesas", "D1", "Purpose")
	f.SetCellValue("Despesas", "E1", "Value")

	// Highlight blank cells
	blankFormat, err := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#FFFF00"},
			Pattern: 1,
		},
	})
	if err != nil {
		return err
	}

	// Set value of cells
	expenseRecordsRow := 1
	for _, record := range expenseRecords {
		expenseRecordsRow++
		f.SetCellValue("Despesas", "A"+strconv.Itoa(expenseRecordsRow), record.Date)
		f.SetCellValue("Despesas", "B"+strconv.Itoa(expenseRecordsRow), record.Company)
		f.SetCellValue("Despesas", "C"+strconv.Itoa(expenseRecordsRow), record.Category)
		f.SetCellValue("Despesas", "D"+strconv.Itoa(expenseRecordsRow), record.Purpose)
		f.SetCellValue("Despesas", "E"+strconv.Itoa(expenseRecordsRow), record.Value)

		// Find out how to use SetConditionalStyle for blank cells
		if record.Company == "" {
			err = f.SetCellStyle("Despesas", "B"+strconv.Itoa(expenseRecordsRow), "B"+strconv.Itoa(expenseRecordsRow), blankFormat)
			if err != nil {
				return err
			}
		}
	}

	classifyManuallyRow := expenseRecordsRow
	for _, record := range classifyManually {
		classifyManuallyRow++
		f.SetCellValue("Despesas", "A"+strconv.Itoa(classifyManuallyRow), record.Date)
		f.SetCellValue("Despesas", "B"+strconv.Itoa(classifyManuallyRow), record.Company)
		f.SetCellValue("Despesas", "C"+strconv.Itoa(classifyManuallyRow), record.Category)
		f.SetCellValue("Despesas", "D"+strconv.Itoa(classifyManuallyRow), record.Purpose)
		f.SetCellValue("Despesas", "E"+strconv.Itoa(classifyManuallyRow), record.Value)

		// Find out how to use SetConditionalStyle for blank cells
		if record.Category == "" {
			err = f.SetCellStyle("Despesas", "C"+strconv.Itoa(classifyManuallyRow), "C"+strconv.Itoa(classifyManuallyRow), blankFormat)
			if err != nil {
				return err
			}
		}

		if record.Purpose == "" {
			err = f.SetCellStyle("Despesas", "D"+strconv.Itoa(classifyManuallyRow), "D"+strconv.Itoa(classifyManuallyRow), blankFormat)
			if err != nil {
				return err
			}
		}
	}

	err = f.SetCellFormula("Despesas", "E"+strconv.Itoa(classifyManuallyRow+1), "=SUM(E1:E"+strconv.Itoa(classifyManuallyRow)+")")
	if err != nil {
		return err
	}

	// Format date cells
	dateFormat, err := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#C6EFCE"},
			Pattern: 1,
		},
	})
	if err != nil {
		return err
	}

	err = f.SetCellStyle("Despesas", "A1", "A"+strconv.Itoa(classifyManuallyRow), dateFormat)
	if err != nil {
		return err
	}

	// Format value cells
	valueFormat, err := f.NewStyle(&excelize.Style{
		NumFmt: 171,
	})
	if err != nil {
		return err
	}

	err = f.SetCellStyle("Despesas", "E1", "E"+strconv.Itoa(classifyManuallyRow), valueFormat)
	if err != nil {
		return err
	}

	expenseFormat, err := f.NewConditionalStyle(&excelize.Style{
		Font: &excelize.Font{Color: "#FF0000"},
	})
	if err != nil {
		return err
	}

	err = f.SetConditionalFormat("Despesas", "E2:E"+strconv.Itoa(classifyManuallyRow), []excelize.ConditionalFormatOptions{
		{
			Type:     "cell",
			Criteria: "<",
			Value:    "0",
			Format:   expenseFormat,
		},
	})
	if err != nil {
		return err
	}

	// Create table by Category
	err = process.CreateCategoryTable(f, classifyManuallyRow)
	if err != nil {
		return err
	}

	// Add chart
	err = utils.AddPieChart(f)
	if err != nil {
		return err
	}

	// Save spreadsheet by the given path.
	if err := f.SaveAs(month + "_budget.xlsx"); err != nil {
		return err
	}

	return nil
}
