package main

import (
	"fmt"
	"net/http"
	"os"
	storage "github.com/MSOpenTech/azure-sdk-for-go/storage"
	fb "github.com/huandu/facebook"
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
    
    setpHttpServer();

    res, _ := fb.Get("/4", fb.Params{
        "fields": "username",
    })
    fmt.Println("here is my facebook username:", res["username"])

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

	httpListenAndServe()
}



func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, page is %s!", r.URL.Path[1:])
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Json page is %s!", r.URL.Path[1:])
}

func setpHttpServer() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/json/", jsonHandler)
}

func httpListenAndServe() {
	http.ListenAndServe(":8080", nil)
}
