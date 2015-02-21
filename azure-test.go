package main

import (
	"fmt"
	"os"
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

}

