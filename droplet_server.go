package main

import (
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
	storage "github.com/MSOpenTech/azure-sdk-for-go/storage"
	fb "github.com/huandu/facebook"
	"code.google.com/p/gorest" 
)

func main() {

    gorest.RegisterService(new(DropletServer)) 
    http.Handle("/",gorest.Handle())     
    http.ListenAndServe(":8080",nil) 
    
    res, _ := fb.Get("/4", fb.Params{
        "fields": "username",
    })
    fmt.Println("here is my facebook username:", res["username"])

}

type Droplet struct{
    IntId             int
    StringName        string
    BoolValue         bool
}

type DropletServer struct { 
    gorest.RestService `root:"/" consumes:"application/json" produces:"application/json"`     

	item    gorest.EndPoint `method:"GET" path:"/item/{Id:int}" output:"Droplet"`
    items   gorest.EndPoint `method:"GET" path:"/items/" output:"[]Droplet"`
    insert  gorest.EndPoint `method:"POST" path:"/insert/" postdata:"[]Droplet"`
    getDroplet    gorest.EndPoint `method:"GET" path:"/d/{userName:string}/{dropletName:string}" output:"Droplet"`
}


func(serv DropletServer) Item(Id int) Droplet {
    serv.ResponseBuilder().SetResponseCode(200)
    item := Droplet {IntId:Id, StringName:"Name with id returned", BoolValue:true}
    return item
}

func(serv DropletServer) Items() []Droplet{
    serv.ResponseBuilder().SetResponseCode(200)
    slice := []Droplet{
      Droplet {IntId:0, StringName:"Name 0", BoolValue:true},
      Droplet {IntId:1, StringName:"Name 1", BoolValue:true},
    }

    item := Droplet {IntId:200, StringName:"Name 4", BoolValue:true}
    slice = append(slice, item)

    return slice
}

func(serv DropletServer) Insert(items []Droplet) {
    fmt.Println("Got a request to insert items")
    fmt.Println("Item Count", len(items))
    serv.ResponseBuilder().SetResponseCode(200)
}

var blobClient *storage.BlobStorageClient

func(serv DropletServer) GetDroplet(userName string, dropletName string) Droplet {
    serv.ResponseBuilder().SetResponseCode(200)

    if blobClient == nil {
    	err := initializeBlobClient()
    	 if err != nil {
    		return Droplet {0,"Blob Client was null",false}
    	}
    }

    resp, err := blobClient.GetBlob(userName, dropletName)
	if err != nil {
		fmt.Printf("Error trying to check blob contents!\n")
	}

	// Verify contents
	respBody, err := ioutil.ReadAll(resp)
	defer resp.Close()
	if err != nil {
		fmt.Printf("Error trying to get blob contents!\n")
	}

	contents := string(respBody[:len(respBody)])

    item := Droplet {IntId:100, StringName:contents, BoolValue:true}
    return item
}


func initializeBlobClient() (error) {
	name := os.Getenv("STORAGE_ACCOUNT_NAME")
	key := os.Getenv("STORAGE_KEY")
	cli, err := storage.NewBasicClient(name, key)

	if err != nil {
		return err		
	}

	blobClient = cli.GetBlobService()
	return nil
}