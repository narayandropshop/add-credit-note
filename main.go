package uploadbills

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	godrej "main.go/godrej"
	marico "main.go/marico"
	"main.go/utils"
)

// gcloud functions deploy UploadCreditNoteAPI  --runtime go113  --trigger-http  --allow-unauthenticated --timeout 540s

func UploadCreditNoteAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// Set CORS headers for the main request.
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var d utils.InputData
	err := json.NewDecoder(r.Body).Decode(&d)

	res := utils.RestStatus{
		Message: "Success",
		Code:    200,
	}

	if err != nil {
		// log.Printf("error parsing application/json: %v", err)
		res = utils.RestStatus{
			Message: "error parsing application/json",
			Code:    500,
		}
		fmt.Fprintf(w, "%+v", res)
		return
	}

	if d.Object == "" {
		fmt.Fprint(w, "File path with name is required!")
		return
	}

	fmt.Println(d)
	startTime := time.Now()

	fmt.Println("Start: ", startTime.String())

	if d.CustomerId == "jdEMMQbw29yhULLeRgJd" {
		godrej.ReadFile(d, "", &res)
	} else if d.CustomerId == "7MI05trvcFT8aez5C1iU" {
		marico.ReadFile(d, "", &res)
	} else {
		res = utils.RestStatus{
			Message: "Not supported CUSTOMER / BRAND type ",
			Code:    500,
		}
		// fmt.Fprintf(w, "%+v", res)
	}

	endTime := time.Now()
	fmt.Println("End: ", endTime.String())
	fmt.Println("Time taken : ", endTime.Sub(startTime))

	resR, _ := json.Marshal(res)

	// fmt.Fprintf(w, "%+v", string(resR))

	w.Header().Set("Content-Type", "application/json")
	w.Write(resR)
	w.WriteHeader(res.Code)
	return
}
