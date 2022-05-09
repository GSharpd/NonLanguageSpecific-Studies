package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Name  string `json:"name"`
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
}
