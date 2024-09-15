package structs

import "time"

type XLSXContent struct {
	TotalInvestment     float64
	TotalPortfolioValue float64
	PortfolioXIRR       float64
	UserDetails         *UserDetails
	SchemeDetails       map[string]*SchemeDetail
}

type SchemeDetail struct {
	SchemeName         string
	AMCName            string
	Category           string
	FolioNumbers       []int64
	InvestedValue      float64
	CurrentValue       float64
	Returns            float64
	Units              float64
	XIRR               float64
	Active             bool
	TransactionDetails []*TransactionDetail
}

type TransactionDetail struct {
	TransactionDescription string
	TransactionDate        time.Time
	NAV                    float64
	Units                  float64
	Amount                 float64
}

type UserDetails struct {
	Name         string
	MobileNumber string
	Email        string
	PAN          string
	FromDate     time.Time
	ToDate       time.Time
}
