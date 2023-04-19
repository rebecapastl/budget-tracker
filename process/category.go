package process

import (
	"strconv"

	"github.com/xuri/excelize"
)

// How to filter only negative values (to sum expenses)
func CreateCategoryTable(f *excelize.File, lastRow int) error {
	err := f.AddPivotTable(&excelize.PivotTableOptions{
		DataRange:       "Despesas!$A$1:$E$" + strconv.Itoa(lastRow),
		PivotTableRange: "Despesas!$G$20:$M$50",
		Rows: []excelize.PivotTableField{
			{Data: "Category", DefaultSubtotal: true}},
		Data: []excelize.PivotTableField{
			{Data: "Value", Name: "Categories Summary"}},
		RowGrandTotals: true,
		ColGrandTotals: true,
		ShowDrill:      true,
		ShowRowHeaders: true,
		ShowColHeaders: true,
		ShowLastColumn: true,
	})
	if err != nil {
		return err
	}

	return nil
}
