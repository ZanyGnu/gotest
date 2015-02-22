package main 

import (
	"fmt"
	"encoding/json"
)

type Message struct {
    Name string
    Body string
    Time int64
}

func main() {

	// create a Message object
	m := Message{"Foo", "Bar", 1294706395881547000}

	// convert the object to json reoresented bytes
	jsonObj, err := json.Marshal(m)

	if err != nil {
		return 
	}

	// decode the json object back to a Type
	var message Message

	err = json.Unmarshal(jsonObj, &message)

	if err != nil {
		return
	}

	fmt.Println(
		"Message has Name ", message.Name,
		" Body = ", message.Body,
		" Time = ", message.Time)
}