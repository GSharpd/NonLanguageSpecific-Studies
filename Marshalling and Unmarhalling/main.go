package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Name  string `json:"name"` // You can use a "-" in order to hide a property during marshalling
	Breed string `json:"breed"`
	Age   uint8  `json:"age"`
}

func main() {
	dog1 := dog{"Rex", "Dalmatian", 3}

	JSONDog1, err := json.Marshal(dog1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(JSONDog1)
	fmt.Println(bytes.NewBuffer(JSONDog1))

	c2 := map[string]string{
		"name":  "Toby",
		"breed": "Poodle",
	}

	JSONDog2, err := json.Marshal(c2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(JSONDog2)
	fmt.Println(bytes.NewBuffer(JSONDog2))

	// Unmarshalling

	dog3 := `{"name":"Max", "breed":"Rotweiller", "age": 7}`

	var dog31 dog

	fmt.Println(dog3)
	fmt.Println(dog31)

	if err := json.Unmarshal([]byte(dog3), &dog31); err != nil {
		log.Fatal(err)
	}

	fmt.Println(dog31)

	dog32 := make(map[string]string)
	dog33 := `{"name":"Max", "breed":"Rotweiller", "age": "7"}`

	if err := json.Unmarshal([]byte(dog33), &dog32); err != nil {
		log.Fatal(err)
	}

	fmt.Println(dog32)
}
