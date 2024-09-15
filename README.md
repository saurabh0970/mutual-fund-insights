# Mutual Fund Insights

A Go CLI that reads your mutual fund portfolio from an Excel file downloaded from [MFCentral](https://www.mfcentral.com/) and provides key insights, including XIRR calculations for individual mutual funds and for the entire portfolio.

## How to Use:

1. **Download your Consolidated Account Statement (CAS)** from MFCentral in Excel format. Ensure that you download the detailed report (including the transaction listing).

2. **Clone this repository** into your local environment. The repo includes a precompiled binary (`mutual-fund-insights`). To directly run this binary, use the command:

    ```bash
    ./mutual-fund-insights run YOUR-FILE-PATH
    ```

   Alternatively, you can run the tool directly without specifying a file path, but in that case, make sure that the CAS report is located in the same directory as the code and is named `report.xlsx`. Use:

    ```bash
    ./mutual-fund-insights run
    ```

3. **Compile the code** yourself by running the provided shell script:

    ```bash
    sh run.sh
    ```

4. **Install the tool globally** using `go install`. This allows you to run the script from anywhere on your machine. Make sure to provide the path to your report file when using the installed version. Run the following command:

    ```bash
    mutual-fund-insights run YOUR-FILE-PATH
    ```

## Requirements

- Go 1.23 or higher
- Excel file from MFCentral (CAS in detailed report format)

## Features

- **XIRR Calculation:** Get XIRR for each individual mutual fund in your portfolio.
- **Portfolio Analysis:** Obtain insights into the entire portfolio’s performance.
