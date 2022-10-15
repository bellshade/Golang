package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name        string `json:"Name"`
	Age         int    `json:"Age,omitempty"`
	PhoneNumber string `json:",omitempty"`
	Address     string `json:"-"`
	Email       string `json:"-,"`
}

func main() {
	person := Person{
		Name:        "Someone",
		Age:         12,
		PhoneNumber: "123456678",
		Address:     "Somestreet",
		Email:       "Someone@mail.com",
	}

	// marshal
	res, err := json.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}
	resString := string(res)
	// hasil setelah di-marshall adalah sebagai berikut
	// {"Name":"Someone","Age":12,"PhoneNumber":"123456678","-":"Someone@mail.com"}
	fmt.Println(resString)

	// unmarshal
	var personUnmarshall Person
	err = json.Unmarshal(res, &personUnmarshall)
	if err != nil {
		log.Fatal(err)
	}
	// hasil setelah di-unmarshall
	fmt.Printf("%v \n", personUnmarshall.Name)
	fmt.Printf("%v \n", personUnmarshall.Age)
	fmt.Printf("%v \n", personUnmarshall.PhoneNumber)
	// atribut Address tidak akan ditampilkan, karena diabaikan pada package ini
	fmt.Printf("%v \n", personUnmarshall.Address)
	fmt.Printf("%v \n", personUnmarshall.Email)
}
