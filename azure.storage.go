package main

import (
	"fmt"
	"os"
	"bytes"
	"crypto/rand"
	storage "github.com/MSOpenTech/azure-sdk-for-go/storage"
)

func getBlobClient() (*storage.BlobStorageClient, error) {
	name := os.Getenv("STORAGE_ACCOUNT_NAME")
	key := os.Getenv("STORAGE_KEY")
	cli, err := storage.NewBasicClient(name, key)

	if err != nil {
		return nil, err
	}

	return cli.GetBlobService(), nil
}


func main() {
    cli, err := getBlobClient()

   	if err != nil {
		fmt.Printf("Error trying to check if container exists!\n")
		return
	}

	cnt := "testcontainer"
	ok, err := cli.ContainerExists(cnt)
	if err != nil {
		fmt.Printf("Error trying to check if container exists!\n")
	}

	if ok {
		fmt.Printf("Found container!\n")
	} else {
		fmt.Printf("Containr not found");
	}	

	blob_operations(cli);
}

func blob_operations(cli *storage.BlobStorageClient) {
	cnt := "testcontainer"
	blob := "testblob-" + randString(10);
	body := []byte(randString(1024))

	// Create the blob and add contents to it
	err := cli.PutBlockBlob(cnt, blob, bytes.NewReader(body))

	// Validate that the blob exists
	ok, err := cli.BlobExists(cnt, blob)
	if err != nil {
		fmt.Printf("Error trying to check if container exists!\n")
	}

	if ok {
		fmt.Printf("Found blob!\n")
	} else {
		fmt.Printf("Blob not found");
	}
}

func randString(n int) string {
	if n <= 0 {
		panic("negative number")
	}
	const alphanum = "0123456789abcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}
