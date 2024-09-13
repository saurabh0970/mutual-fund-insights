package implementations

import (
	"github.com/maksim77/goxirr"
	"scrpits/process"
	"scrpits/structs"
	"scrpits/validator"
	"time"
)

func CalculateXIRR(fileName string) (*structs.XLSXContent, error) {
	xlsxContent, err := process.GetXLSXContentFromFilename(fileName)
	validator.Must(err)

	allTransactions := make([]goxirr.Transaction, 0)

	for _, schemeDetail := range xlsxContent.SchemeDetails {
		currTransactions := make([]goxirr.Transaction, len(schemeDetail.TransactionDetails)+1)
		for i, transactionDetail := range schemeDetail.TransactionDetails {
			schemeDetail.InvestedValue += transactionDetail.Amount
			currTransactions[i] = goxirr.Transaction{
				Date: transactionDetail.TransactionDate,
				Cash: transactionDetail.Amount,
			}
		}
		// Adding because invested value is taken as negative
		schemeDetail.Returns = schemeDetail.CurrentValue + schemeDetail.InvestedValue
		if schemeDetail.Active {
			currTransactions[len(schemeDetail.TransactionDetails)] = goxirr.Transaction{
				Date: time.Now(),
				Cash: schemeDetail.CurrentValue,
			}
		}
		schemeDetail.XIRR = goxirr.Xirr(currTransactions)

		allTransactions = append(allTransactions, currTransactions...)
	}
	xlsxContent.PortfolioXIRR = goxirr.Xirr(allTransactions)
	return xlsxContent, nil
}
