package cmd

import (
	"github.com/spf13/cobra"
	"mutual-fund-insights/implementations"
	"mutual-fund-insights/validator"
)

// runCmd runs the script for the selected portfolio
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the script for selected portfolio",
	Run: func(cmd *cobra.Command, args []string) {
		filePath := "report.xlsx"
		if len(args) == 1 {
			filePath = args[0]
		}
		xlsxContent, err := implementations.CalculateXIRR(filePath)
		validator.Must(err)
		implementations.AnalyzePortfolio(xlsxContent)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
