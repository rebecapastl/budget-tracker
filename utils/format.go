package utils

import (
	"fmt"
	"strings"
)

func FormatDate(oldDate string) string {
	// format string
	slicedDate := strings.Split(oldDate, "/")
	stringDate := fmt.Sprintf("%04s-%02s-%02s", slicedDate[2], slicedDate[0], slicedDate[1])

	return stringDate
}

func SanitizeCompanyName(input string) string {
	company := input

	if strings.Contains(strings.ToUpper(input), "AMAZON.") {
		company = "Amazon"
	} else if strings.Contains(input, "PC EXPRESS") {
		company = "Canadian Superstore"
	} else if strings.Contains(input, "BUY LOW FOODS") {
		company = "Buy Low Foods"
	} else if strings.Contains(input, "DOLLAR TREE") {
		company = "Dollar Tree"
	} else if strings.Contains(input, "DOLLAR PLUS SHOP") {
		company = "Dollar Shop"
	} else if strings.Contains(input, "UBER") {
		company = "Uber"
	} else if strings.Contains(input, "LYFT") {
		company = "Lyft"
	} else if strings.Contains(input, "SKIPTHEDISHES") {
		company = "Skip The Dishes"
	} else if strings.Contains(input, "SHOPPERS DRUG MART") {
		company = "Shoppers Drug Mart"
	} else if strings.Contains(input, "LONDON DRUGS") {
		company = "London Drugs"
	} else if strings.Contains(input, "TRU EARTH") {
		company = "Tru Earth"
	} else if strings.Contains(input, "CRAVE") {
		company = "Crave"
	} else if strings.Contains(input, "ROGERS") {
		company = "Rogers"
	} else if strings.Contains(input, "COMPASS") {
		company = "Compass Card"
	} else if strings.Contains(input, "REVENUE SERVICES BC") {
		company = "Revenue Services BC"
	} else if strings.Contains(input, "CANADA") || strings.Contains(input, "SOCIAL DEV") {
		company = "Canadian Government"
	} else if strings.Contains(input, "FROM - *****") {
		company = "Credit Card"
	}

	return company
}
