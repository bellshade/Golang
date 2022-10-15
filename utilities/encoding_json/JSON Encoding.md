# Tentang JSON Encoding

## Encoding & Decoding
Dalam dunia komputer, *encoding* adalah proses menerjemahkan suatu informasi ke format tertentu, sedangkan *decoding* adalah proses yang berlawanan dengan *encoding*, yakni mengubah suatu informasi dari format tertentu ke bentuk aslinya.

## Marshal & Unmarshal JSON
Sederhananya, *marshal* dan *unmarshal* memiliki kesamaan dengan *encoding* dan *decoding*, yaitu menerjemahkan informasi ke format tertentu. Namun, biasanya istilah *marshal* dan *unmarshal* identik dengan menerjemahkan informasi ke dalam format *JSON* (Javascript Object Notation).

## Marshal & Unmarshal JSON Di Golang
Dalam bahasa pemrograman `Go`, *JSON* biasanya digunakan untuk *file* konfigurasi atau *request* dan *response* dari HTTP.

### Marshaling
Ada beberapa aturan dari `Go` ketika kita ingin melakukan *marshal* dari `struct` ke *JSON* sebagai berikut.

```
// Atribut Field akan muncul sebagai key "myName" dalam format JSON.
Field int `json:"myName"`

// Atribut Field tidak akan muncul sebagai key "myName" pada
// format JSON jika nilainya tidak ada dan sebaliknya
Field int `json:"myName,omitempty"`

// Secara default, atribut FIeld akan menjadi key "Field" di JSON
// namun akan dikosongkan jika atribut Field tidak ada nilainya.
Field int `json:",omitempty"`

// Atribut Field diabaikan dari package ini
Field int `json:"-"`

// Atribut Field akan muncul sebagai key "-" pada JSON
Field int `json:"-,"`
```

Berikut contoh sederhana dari proses *marshal* `struct` ke *JSON*.

```Go
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
	res, err := json.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}
	var personUnmarshall Person
	err = json.Unmarshal(res, &personUnmarshall)
	if err != nil {
		log.Fatal(err)
	}
	resString := string(res)
	// hasil setelah di-marshall adalah sebagai berikut
	// {"Name":"Someone","Age":12,"PhoneNumber":"123456678","-":"Someone@mail.com"}
	fmt.Println(resString)
}
```