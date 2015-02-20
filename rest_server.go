package main 

import ( 
      "code.google.com/p/gorest" 
      "net/http" 
      "fmt"
)

func main() { 
    // GoREST usage: http://localhost:8181/tutorial/hello 
    gorest.RegisterService(new(Tutorial)) //Register our service 
    http.Handle("/",gorest.Handle())     
    http.ListenAndServe(":8181",nil) 
}

//Service Definition 
type Tutorial struct { 
    gorest.RestService `root:"/tutorial/" consumes:"application/json" produces:"application/json"` 
    hello  gorest.EndPoint `method:"GET" path:"/hello/" output:"string"` 
    insert   gorest.EndPoint `method:"POST" path:"/insert/" postdata:"int"`
}

func(serv Tutorial) Hello() string{ 
    return "Hello World" 
}

func(serv Tutorial) Insert(number int) { 
    fmt.Println("Got a request to insert ")
    serv.ResponseBuilder().SetResponseCode(200) 
}