package process

import (
	"errors"
	"fmt"
	"github.com/araddon/dateparse"
	"github.com/xuri/excelize/v2"
	"mutual-fund-insights/consts"
	"mutual-fund-insights/structs"
	"mutual-fund-insights/validator"
	"strconv"
)

func GetXLSXContentFromFilename(fileName string) (resp *structs.XLSXContent, err error) {
	resp = &structs.XLSXContent{}
	excelFile, err := excelize.OpenFile(fileName)
	validator.Must(err)
	defer func() {
		if err := excelFile.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	if excelFile.SheetCount != consts.SheetCount {
		validator.Must(errors.New("excel file has wrong number of sheets"))
	}

	schemeNameToSchemeDetailsMap := make(map[string]*structs.SchemeDetail)

	portfolioDetailsSheet, err := excelFile.GetRows(consts.SheetPortfolioDetails)
	validator.Must(err)

	totalPortfolioValue, userDetails := fillSchemeDetails(portfolioDetailsSheet, schemeNameToSchemeDetailsMap)

	transactionDetailsSheet, err := excelFile.GetRows(consts.SheetTransactionDetails)
	validator.Must(err)
	totalInvestment := fillTransactionDetails(transactionDetailsSheet, schemeNameToSchemeDetailsMap)

	resp.UserDetails = userDetails
	resp.TotalInvestment = totalInvestment
	resp.SchemeDetails = schemeNameToSchemeDetailsMap
	resp.TotalPortfolioValue = totalPortfolioValue
	return resp, nil
}

func fillTransactionDetails(transactionDetailsSheet [][]string, schemeNameToSchemeDetailsMap map[string]*structs.SchemeDetail) float64 {
	reachedTransactionRows := false
	totalInvestment := 0.0
	for _, val := range transactionDetailsSheet {
		if reachedTransactionRows {
			transTime, err := dateparse.ParseAny(val[2])
			validator.Must(err)
			nav, err := strconv.ParseFloat(val[3], 64)
			validator.Must(err)
			units, err := strconv.ParseFloat(val[4], 64)
			validator.Must(err)
			amount, err := strconv.ParseFloat(val[5], 64)
			validator.Must(err)
			totalInvestment += amount
			// Setting amount negative when invested and positive when redeemed
			schemeNameToSchemeDetailsMap[val[0]].TransactionDetails = append(schemeNameToSchemeDetailsMap[val[0]].TransactionDetails, &structs.TransactionDetail{
				TransactionDescription: val[1],
				TransactionDate:        transTime,
				NAV:                    nav,
				Units:                  units,
				Amount:                 -amount,
			})
		} else if len(val) != 0 {
			if val[0] == consts.ColNameSchemeName {
				reachedTransactionRows = true
			}
		}
	}
	return totalInvestment
}

func fillSchemeDetails(portfolioDetailsSheet [][]string, schemeNameToSchemeDetailsMap map[string]*structs.SchemeDetail) (float64, *structs.UserDetails) {
	reachedSchemeDetailRows := false
	totalPortfolioValue := -1.0
	userDetails := &structs.UserDetails{
		Name: "Investor",
	}
	for i, val := range portfolioDetailsSheet {
		if reachedSchemeDetailRows {
			currFolioNum, err := strconv.ParseInt(val[3], 10, 64)
			validator.Must(err)
			currVal, err := strconv.ParseFloat(val[5], 64)
			validator.Must(err)
			currUnits, err := strconv.ParseFloat(val[7], 64)
			validator.Must(err)
			schemeActive := false
			if currUnits > 0 {
				schemeActive = true
			}
			if schemeDetail, ok := schemeNameToSchemeDetailsMap[val[0]]; ok {
				schemeDetail.FolioNumbers = append(schemeDetail.FolioNumbers, currFolioNum)
				schemeDetail.CurrentValue += currVal
				schemeDetail.Units += currUnits
			} else {
				schemeNameToSchemeDetailsMap[val[0]] = &structs.SchemeDetail{
					SchemeName:         val[0],
					AMCName:            val[1],
					Category:           val[2],
					FolioNumbers:       []int64{currFolioNum},
					CurrentValue:       currVal,
					Units:              currUnits,
					Active:             schemeActive,
					TransactionDetails: make([]*structs.TransactionDetail, 0),
				}
			}
		} else if len(val) >= 1 && val[0] == consts.ColNameSchemeName {
			reachedSchemeDetailRows = true
		} else if totalPortfolioValue == -1 && validateForTotalPortfolioValue(val, portfolioDetailsSheet, i) {
			totalPortValue, err := strconv.ParseFloat(portfolioDetailsSheet[i+1][1], 64)
			validator.Must(err)
			totalPortfolioValue = totalPortValue
		} else if len(val) >= 2 && val[0] == consts.LabelToDate {
			toDate, err := dateparse.ParseAny(val[1])
			validator.Must(err)
			userDetails.ToDate = toDate
		} else if len(val) >= 2 && val[0] == consts.LabelFromDate {
			fromDate, err := dateparse.ParseAny(val[1])
			validator.Must(err)
			userDetails.FromDate = fromDate
		} else if len(val) >= 2 && val[0] == consts.LabelPAN {
			userDetails.PAN = val[1]
		} else if len(val) >= 2 && val[0] == consts.LabelEmail {
			userDetails.Email = val[1]
		} else if len(val) >= 2 && val[0] == consts.LabelMobileNumber {
			userDetails.MobileNumber = val[1]
		} else if len(val) >= 2 && val[0] == consts.LabelName {
			userDetails.Name = val[1]
		}
	}
	return totalPortfolioValue, userDetails
}

func validateForTotalPortfolioValue(val []string, portfolioDetailsSheet [][]string, i int) bool {
	return len(val) >= 2 && val[1] == consts.ColCurrentPortfolioValue && len(portfolioDetailsSheet) > i+1 && len(portfolioDetailsSheet[i+1]) >= 2
}
