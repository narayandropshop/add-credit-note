package upload

import (
	"fmt"
	"os"
	"sync"

	"google.golang.org/api/option"
	"main.go/utils"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"golang.org/x/net/context"
)

var client *firestore.Client
var ctx context.Context

const THREADCOUNT = 1000

func createClient() {

}

/*
UploadBills:
Create client object
Map consumer bill ids with bills
Send data to upload bills
*/
func UploadBills(dataInput utils.InputData, billsIdsData map[string][]utils.Bills, projectID string, resStatus *utils.RestStatus) {
	ctx = context.Background()

	count := 0
	var err error
	var app *firebase.App

	if os.Getenv("TEST") != "true" {
		conf := &firebase.Config{ProjectID: projectID}
		app, err = firebase.NewApp(ctx, conf)
	} else {
		sa := option.WithCredentialsFile("/home/stingray/dropshop/tools/dropshop-5cbbf-ecc067181e26.json")
		app, err = firebase.NewApp(ctx, nil, sa)
	}

	if err != nil {
		// log.Fatalln(err)
		panic("An error occurred (panic)")

	}

	client, err = app.Firestore(ctx)
	if err != nil {
		// log.Fatalln(err)
		panic("An error occurred (panic)")

	}

	defer client.Close()

	processBills(billsIdsData, resStatus, dataInput)
	fmt.Println("processed ", count)
}

func processBills(billsIdsData map[string][]utils.Bills, resStatus *utils.RestStatus, dataInput utils.InputData) {
	noOfBillsIds := len(billsIdsData)
	ch := make(chan []utils.Bills, noOfBillsIds)

	for _, billsData := range billsIdsData {
		ch <- billsData
	}

	close(ch)
	wg := &sync.WaitGroup{}
	wg.Add(THREADCOUNT)
	for g := 0; g < THREADCOUNT; g++ {
		go func(resStatus *utils.RestStatus) {
			defer wg.Done()
			idx := 0
			for u := range ch {
				utils.Process(idx, u, ctx, client, resStatus, dataInput)
				idx++
			}
		}(resStatus)
	}

	wg.Wait()
}
