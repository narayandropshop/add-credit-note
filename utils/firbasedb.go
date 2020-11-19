package utils

import (
	"log"
	"sync"

	"cloud.google.com/go/firestore"
	cmap "github.com/orcaman/concurrent-map"
	"golang.org/x/net/context"
)

var customer_id = "WnwBuNNbGQLHNRAOStli"
var warehouse_id = "XcSWkvuxvd4ehu1YJDbB"
var store_type = "TT"

var client *firestore.Client
var ctx context.Context

// var docObj *firestore.DocumentRef
var mu sync.Mutex
var dataInput InputData

// var docObj *firestore.Client

var billIdMap cmap.ConcurrentMap
var THREADCOUNT = 1000

/*
Process :
Check for existence of the data, create, update and upload bills
*/
func Process(counter int, data []Bills, ctxObj context.Context, cliObj *firestore.Client, resStatus *RestStatus, dataInput InputData) {
	setDbObj(ctxObj, cliObj, dataInput)

	for _, bill := range data {
		bill_exists := isBillExists(&bill)

		if !bill_exists {
			mu.Lock()
			resStatus.ExistsBillCount += 1
			mu.Unlock()
			return
		}

		updateBillData(&bill)
	}
}

func setDbObj(ctxObj context.Context, cliObj *firestore.Client, input InputData) {
	ctx = ctxObj
	client = cliObj
	dataInput = input
}

func getCollectionObj(key string) *firestore.CollectionRef {
	if dataInput.Prod == true {
		return client.Collection(key)
	} else {
		return client.Collection("debug").Doc(dataInput.Table).Collection(key)
	}
}

/*
isBillExists:
Check if bill is exists
*/
func isBillExists(bill *Bills) bool {
	collectionObj := getCollectionObj("bills")
	// docObj.Collection("bills")
	bill_id := ""
	isExist := false
	bill_query := collectionObj.Where("customerBillId", "==", bill.CustomerBillID).Documents(ctx)

	query, _ := bill_query.GetAll()

	if len(query) > 0 {
		for _, q := range query {
			bill_id = q.Data()["billId"].(string)
		}
		isExist = true
	} else {
		isExist = false
	}

	bill.BillID = bill_id
	return isExist
}

/*
updateBillData:
Update bill info
*/
func updateBillData(bill *Bills) {
	collectionObj := getCollectionObj("bills").Doc(bill.BillID)
	itemInfo := map[string]interface{}{
		"creditNote": bill.CreditNote,
	}

	_, err := collectionObj.Set(ctx, itemInfo, firestore.MergeAll)
	logError(err)
}

func logError(err error) {
	if err != nil {
		log.Println(err)
		panic("An error occurred (panic)")

	}
}
