package main 

import ( 
      "code.google.com/p/gorest" 
      "net/http" 
      "fmt"
)

type Item struct{
    IntId             int
    StringName        string
    BoolValue         bool
}

func main() { 
    // GoREST usage: http://localhost:8181/tutorial/hello 
    gorest.RegisterService(new(Tutorial)) //Register our service 
    http.Handle("/",gorest.Handle())     
    http.ListenAndServe(":8181",nil) 
}

//Service Definition 
type Tutorial struct { 
    gorest.RestService `root:"/tutorial/" consumes:"application/json" produces:"application/json"` 
    hello   gorest.EndPoint `method:"GET" path:"/hello/" output:"string"`

    items   gorest.EndPoint `method:"GET" path:"/items/" output:"[]Item"`
    insert  gorest.EndPoint `method:"POST" path:"/insert/" postdata:"[]Item"`

    insertint  gorest.EndPoint `method:"POST" path:"/insertint/" postdata:"int"`
}

func(serv Tutorial) Hello() string{ 
    return "Hello World" 
}

func(serv Tutorial) Items() []Item{
    serv.ResponseBuilder().SetResponseCode(200)
    slice := []Item{
      Item {IntId:0, StringName:"Name 0", BoolValue:true},
      Item {IntId:1, StringName:"Name 1", BoolValue:true},
    }

    item := Item {IntId:200, StringName:"Name 4", BoolValue:true}
    slice = append(slice, item)

    return slice
}

func(serv Tutorial) Insert(items []Item) {
    fmt.Println("Got a request to insert items")
    fmt.Println("Item Count", len(items))
    serv.ResponseBuilder().SetResponseCode(200)
}


func(serv Tutorial) Insertint(number int) {
    fmt.Println("Got a request to insert an int ")
    serv.ResponseBuilder().SetResponseCode(200)
}

 