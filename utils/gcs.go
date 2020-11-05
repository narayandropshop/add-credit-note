package utils

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"time"

	"cloud.google.com/go/storage"
)

func UploadFile(path, name string, data []byte) {
	// Prevent log from printing out time information
	log.SetFlags(0)

	var bucket string
	bucket = "bills_upload"
	// source = "/home/stingray/Downloads/RB-TT.xlsx"
	name = name + ".xlsx"

	var r io.Reader

	r = bytes.NewReader(data)
	ctx := context.Background()
	// sa := option.WithCredentialsFile("/home/stingray/dropshop/tools/dropshop-5cbbf-ecc067181e26.json")

	client, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Println("error in creating clinet")
		log.Printf("storage.NewClient: %v", err)
	}
	defer client.Close()

	object := path + "/" + name

	// Upload an object with storage.Writer.
	wc := client.Bucket(bucket).Object(object).NewWriter(ctx)
	if _, err = io.Copy(wc, r); err != nil {
		log.Printf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		log.Printf("Writer.Close: %v", err)
	}

}

func uploadFile(bucket, object string) {

}

func DownloadFile(bucket, object string) ([]byte, error) {
	// bucket := "bills_upload"
	// object := "RB-TT.xlsx"
	ctx := context.Background()
	// sa := option.WithCredentialsFile("/home/stingray/dropshop/tools/dropshop-5cbbf-ecc067181e26.json")

	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	rc, err := client.Bucket(bucket).Object(object).NewReader(ctx)
	if err != nil {
		return nil, fmt.Errorf("Object(%q).NewReader: %v", object, err)
	}
	defer rc.Close()

	data, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll: %v", err)
	}

	return data, nil
}
