package main

import (
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
	"bytes"
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
    Id             int
    Name           string
    Content        string
}

type DropletServer struct { 
    gorest.RestService `root:"/" consumes:"application/json" produces:"application/json"`     

	item    gorest.EndPoint `method:"GET" path:"/item/{Id:int}" output:"Droplet"`
    items   gorest.EndPoint `method:"GET" path:"/items/" output:"[]Droplet"`
    insert  gorest.EndPoint `method:"POST" path:"/insert/" postdata:"[]Droplet"`
    getDroplet    gorest.EndPoint `method:"GET" 	path:"/d/{userName:string}/{dropletName:string}" output:"Droplet"`
    getDroplets   gorest.EndPoint `method:"GET" 	path:"/d/{userName:string}" output:"[]Droplet"`
    putDroplet    gorest.EndPoint `method:"POST" 	path:"/d/{userName:string}/{dropletName:string}" postdata:"Droplet"`
    putDroplets   gorest.EndPoint `method:"POST" 	path:"/d/{userName:string}" postdata:"[]Droplet"`
}


func(serv DropletServer) Item(Id int) Droplet {
    serv.ResponseBuilder().SetResponseCode(200)
    item := Droplet {Id:Id, Content:"Name with id returned"}
    return item
}

func(serv DropletServer) Items() []Droplet{
    serv.ResponseBuilder().SetResponseCode(200)
    slice := []Droplet{
      Droplet {Id:0, Content:"Name 0"},
      Droplet {Id:1, Content:"Name 1"},
    }

    item := Droplet {Id:200, Content:"Name 4"}
    slice = append(slice, item)

    return slice
}

func(serv DropletServer) Insert(items []Droplet) {
    fmt.Println("Got a request to insert items")
    fmt.Println("Item Count", len(items))
    serv.ResponseBuilder().SetResponseCode(200)
}

var blobClient *storage.BlobStorageClient

func(serv DropletServer) GetDroplets(userName string) []Droplet {
	if blobClient == nil {
    	err := initializeBlobClient()
    	 if err != nil {
    		
    	}
    }

	droplets := []Droplet{}
	marker := ""
	for {
		resp, err := blobClient.ListBlobs(userName, storage.ListBlobsParameters{
			MaxResults: 1024,
			Marker:     marker})
		if err != nil {
			serv.ResponseBuilder().SetResponseCode(500)
			return []Droplet{}
		}

		for _, v := range resp.Blobs {

			fmt.Println("getDroplets: ", v.Name)
			err, droplet := getDropletItem(userName, v.Name)

		    if err != nil {
		    	serv.ResponseBuilder().SetResponseCode(500)
		    	return nil
		    }

			droplets = append(droplets, droplet)
		}

		marker = resp.NextMarker

		if marker == "" || len(resp.Blobs) == 0 {
			break
		}
	}

	return droplets;	
}

func(serv DropletServer) GetDroplet(userName string, dropletName string) Droplet {
    serv.ResponseBuilder().SetResponseCode(200)

    if blobClient == nil {
    	err := initializeBlobClient()
    	 if err != nil {
    	 	serv.ResponseBuilder().SetResponseCode(500)
    		return Droplet{}
    	}
    }

    err, droplet := getDropletItem(userName, dropletName)

    if err != nil {
    	serv.ResponseBuilder().SetResponseCode(500)
    	return Droplet{}
    }

    return droplet
}

func getDropletItem(userName string, dropletName string) (error, Droplet) {

    resp, err := blobClient.GetBlob(userName, dropletName)

	if err != nil {
		fmt.Println("getDropletItem: Error trying to check blob contents!\n")
		return err, Droplet{}
	}

	respBody, err := ioutil.ReadAll(resp)
	
	defer resp.Close()

	if err != nil {
		fmt.Println("getDropletItem: Error trying to get blob contents!\n")
		return err, Droplet{}
	}

	contents := string(respBody[:len(respBody)])

    item := Droplet {Id:100, Name:dropletName, Content:contents}

    return nil, item
}


func(serv DropletServer) PutDroplet(droplet Droplet, userName string, dropletName string)  {
    
    // Create the blob and add contents to it
	err := putDropletItem(userName, droplet.Name, droplet)

	if err == nil {
		serv.ResponseBuilder().SetResponseCode(200)
	} else {
		serv.ResponseBuilder().SetResponseCode(500)
	}
}

func putDropletItem(userName string, dropletName string, droplet Droplet) (error) {
	if blobClient == nil {
    	err := initializeBlobClient()
    	 if err != nil {
    	 	fmt.Println("putDropletItem: ", err)
    		return err;
    	}
    }

    // Create the blob and add contents to it
	return blobClient.PutBlockBlob(userName, dropletName, bytes.NewReader([]byte(droplet.Content)))
}

func(serv DropletServer) PutDroplets(droplets []Droplet, userName string)  {
    
    if blobClient == nil {
    	err := initializeBlobClient()
    	 if err != nil {
    	 	fmt.Println("PutDroplets: ", err)
    		return;
    	}
    }

    for _, droplet := range droplets { 
    	fmt.Println("PutDroplets: Processing ", droplet.Name)
	    err := putDropletItem(userName, droplet.Name, droplet)

		if err != nil {
			fmt.Println("PutDroplets: Error ", err)
			serv.ResponseBuilder().SetResponseCode(500)
			return
		}
	}

	serv.ResponseBuilder().SetResponseCode(200)
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