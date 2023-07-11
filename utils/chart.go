package utils

import (
	"github.com/xuri/excelize/v2"
)

func AddPieChart(f *excelize.File) error {
	err := f.AddChart("Despesas", "F1", &excelize.Chart{
		Type: excelize.Pie,
		Series: []excelize.ChartSeries{
			{
				Name:       "Amount",
				Categories: "Despesas!$G$21:$G$29",
				Values:     "Despesas!$H$21:$H$29",
			},
		},
		Format: excelize.GraphicOptions{
			OffsetX: 15,
			OffsetY: 10,
		},
		Title: excelize.ChartTitle{
			Name: "Category Chart",
		},
		PlotArea: excelize.ChartPlotArea{
			ShowPercent: true,
		},
	})
	if err != nil {
		return err
	}

	return nil
}
