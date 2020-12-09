package utils

import (
	"cloud.google.com/go/firestore"
	"golang.org/x/net/context"
)

type Bills struct {
	BillID              string
	CustomerBillID      string
	CustomerRetailerID  string
	BeatID              string
	Name                string
	GSTIN               string
	BeatName            string
	CustomerProductID   string
	ProductName         string
	Brand               string
	FinalAmount         float64
	GrossAmount         float64
	TotalDiscountAmount float64
	TotalTax            float64
	CGST                float64
	SGST                float64
	Loose               int64
	GST                 float64
	Cases               int64
	PiecesInCase        int64
	QuantityOffered     int64
	QuantityOrdered     int64
	QuantityTotal       int64
	MRP                 float64
	NetAmount           float64
	SellingPrice        float64
	BillDate            string
	RetailerID          string
	ProductID           string
	TotalBillAmount     float64
	TotalBillQuantity   int64
	SkuCount            int64
	CreditNote          float64
	CustomerID          string
	CustomerIDRetailer  string
	DsrId               string
	WarehouseID         string
	StoreType           string
	UploadType          string
	Category            string
	Address             string
	DND                 bool
}

type Retailers struct {
	BeatIds    []string `firestore:"beatIds"`
	RetailerID string   `firestore:"retailerID"`
}

type Beats struct {
	BeatId      string `firestore:"beatId"`
	BeatName    string `firestore:"beatName"`
	CustomerId  string `firestore:"customerId"`
	DsrId       string `firestore:"dsrId"`
	WarehouseId string `firestore:"name"`
}

type FirestoreClient struct {
	Client *firestore.Client
	Ctx    context.Context
}

type InputData struct {
	Object      string `json:"filePath"`
	WarehouseId string `json:"warehouseId"`
	CustomerId  string `json:"customerId"`
	StoreType   string `json:"storeType"`
	Brand       string `json:"brand"`
	Prod        bool   `json:"prod"`
	Table       string `json:"table"`
}

type RestStatus struct {
	Message         string
	Code            int
	BillCount       int
	ExistsBillCount int
}
