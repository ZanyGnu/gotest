package main

import (
	"fmt"
	"net/http"
	"os"
	storage "github.com/MSOpenTech/azure-sdk-for-go/storage"
	fb "github.com/huandu/facebook"
	"code.google.com/p/gorest" 
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

    gorest.RegisterService(new(DropletServer)) 
    http.Handle("/",gorest.Handle())     
    http.ListenAndServe(":8080",nil) 
    
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