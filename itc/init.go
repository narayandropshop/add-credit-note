package itc

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"main.go/upload"
	"main.go/utils"
)

var matchIndexes = []int{0, 21, 20}
var coloumnNames = []string{"InvoiceID", "Adjusted Amount", "Adj Ref"}
var dbIndexs = []int{5}
var sheetStartIndex = 0

/*
ReadFile:
Create bill object from the required build
Build data for upload
*/
func ReadFile(dataInput utils.InputData, projectID string, resStatus *utils.RestStatus) *utils.RestStatus {
	bucket := "bills_upload"
	projectID = "dropshop-5cbbf"

	var fileObj *excelize.File
	var err error
	if os.Getenv("TEST") != "true" {
		fmt.Println("in prod")
		fileData, _ := utils.DownloadFile(bucket, dataInput.Object)
		r := bytes.NewReader(fileData)
		fileObj, err = excelize.OpenReader(r)
	} else {
		fmt.Println("in test")
		fileObj, err = excelize.OpenFile("/home/stingray/Downloads/ksBQR6afWnyaP7G9lv22_creditnote_2021-02-11_Credit Note 11-02-2021_.xlsx")
	}

	if err != nil {
		fmt.Println(err)
		resStatus.Message = "Failed to read file."
		resStatus.Code = 500
		return resStatus
	}

	sheetName := fileObj.GetSheetName(0)
	rows, _ := fileObj.GetRows(sheetName)
	if validateFile(rows[sheetStartIndex]) {
		buildData(dataInput, rows, projectID, resStatus)
	} else {
		resStatus.Message = "Bad xls sheet. OR Please check the brand/cutomer type is correct"
		resStatus.Code = 500
	}

	return resStatus
}

func validateFile(columnsRow []string) bool {
	validateFlag := true
	for i, matchI := range matchIndexes {
		if coloumnNames[i] != columnsRow[matchI] {
			fmt.Println("Bad xls sheet...")
			validateFlag = false
			break
		}
	}

	return validateFlag
}

func buildData(dataInput utils.InputData, rows [][]string, projectID string, resStatus *utils.RestStatus) {
	var dbData []utils.Bills
	biilIdsMap := make(map[string][]utils.Bills)
	count := 0
	for _, row := range rows[sheetStartIndex+1:] {
		var newBill utils.Bills
		if len(row) > 51 {
			if strings.Contains(row[36], "GrandTotal") && !strings.Contains(row[0], "GrandTotal") {
				newBill.UploadType = dataInput.Brand
				newBill.CustomerBillID = row[0]
				newBill.CreditNote = utils.GetFloatFromStirng(row[21])

				if strings.Contains(row[20], "GV") {
					newBill.DND = false
				} else {
					newBill.DND = true
				}

				newBill.CustomerID = dataInput.CustomerId
				newBill.WarehouseID = dataInput.WarehouseId
				newBill.StoreType = dataInput.StoreType

				biilIdsMap[newBill.CustomerBillID] = append(biilIdsMap[newBill.CustomerBillID], newBill)
				dbData = append(dbData, newBill)
				count = count + 1
			}
		} else {
			// if len(row) > 0 {
			// 	resStatus.Code = 500
			// 	resStatus.Message = "Less number of columns"
			// 	break
			// }
		}
	}

	fmt.Println("total entries", count)
	if resStatus.Code != 500 {
		upload.UploadBills(dataInput, biilIdsMap, projectID, resStatus)
		resStatus.BillCount = len(biilIdsMap)
	}
}
