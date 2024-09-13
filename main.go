package main

import (
	"scrpits/implementations"
	"scrpits/validator"
)

func main() {
	xlsxContent, err := implementations.CalculateXIRR("report.xlsx")
	validator.Must(err)
	implementations.AnalyzePortfolio(xlsxContent)
}
