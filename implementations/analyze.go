package implementations

import (
	"fmt"
	"github.com/punithck/indianmoney"
	"mutual-fund-insights/structs"
)

func AnalyzePortfolio(xlsxContent *structs.XLSXContent) {
	fmt.Printf("Hi %v!\n", xlsxContent.UserDetails.Name)
	fmt.Printf("Till date, you have invested ₹%v combined \n", indianmoney.FormatMoneyFloat64(xlsxContent.TotalInvestment, 2))
	fmt.Printf("Your Total Portfolio Value is: ₹%v \n", indianmoney.FormatMoneyFloat64(xlsxContent.TotalPortfolioValue, 2))
	fmt.Printf("Current XIRR for your portfolio is %v%%\n\n", xlsxContent.PortfolioXIRR)

	fmt.Println("The details for your active schemes:")
	for _, schemeDetail := range xlsxContent.SchemeDetails {
		if schemeDetail.Active {
			fmt.Printf("You have invested %v in the scheme %v, the current value is %v, which takes your absolute "+
				"profit to %v and absolute profit percentage to %.2f%%, and the XIRR is %v%%\n", formatMoney(-schemeDetail.InvestedValue),
				schemeDetail.SchemeName, formatMoney(schemeDetail.CurrentValue), formatMoney(schemeDetail.Returns),
				(schemeDetail.Returns/-schemeDetail.InvestedValue)*100, schemeDetail.XIRR)
		}
	}

	fmt.Println("\nThe details for your past schemes:")
	for _, schemeDetail := range xlsxContent.SchemeDetails {
		if !schemeDetail.Active {
			fmt.Printf("For the scheme %v, you made a total profit of %v and the XIRR is %v%%\n",
				schemeDetail.SchemeName, formatMoney(schemeDetail.Returns), schemeDetail.XIRR)
		}
	}
}

func formatMoney(money float64) string {
	return "₹" + indianmoney.FormatMoneyFloat64(money, 2)
}
