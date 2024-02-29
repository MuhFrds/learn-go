package main

import (
	"fmt"
)

func main() {

	// var person map[string]string = map[string]string{}
	// person["name"] = "Eko"
	// person["address"] = "Bandung"

	person := map[string]string{
		"name":    "Eko",
		"address": "Bandung",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Eko"
	book["ups"] = "salah"

	fmt.Println(book)

	delete(book, "ups")

	fmt.Println(book)

}