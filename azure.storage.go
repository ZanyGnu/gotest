package main

import (
	"fmt"
	"os"
	"bytes"
	"crypto/rand"
	"io/ioutil"
	"reflect"
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
	body := []byte(randString(64))

	fmt.Print("Creating blob ", blob,  " and putting contents into it ", body)

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


	// Validate that the contents of the blob are as expected

	resp, err := cli.GetBlob(cnt, blob)
	if err != nil {
		fmt.Printf("Error trying to check blob contents!\n")
	}

	// Verify contents
	respBody, err := ioutil.ReadAll(resp)
	defer resp.Close()
	if err != nil {
		fmt.Printf("Error trying to get blob contents!\n")
	}

	if !reflect.DeepEqual(body, respBody) {
		fmt.Printf("Wrong blob contents.\nExpected: %d bytes, Got: %d byes", len(body), len(respBody))
	}

	fmt.Print("Validated that blob ", blob, " has contents: ", resp)

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
