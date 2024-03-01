package main

import (
	"fmt"
)

func main() {
	name := "Eko"

	switch name{

	case "Eko":
		fmt.Println("Hello Eko")

	case "Budi":
		fmt.Println("Hello Budi")

	default:
		fmt.Println("Hi, Boleh Kenalan? ")
	}

	// Short statement
	switch length := len(name); length > 5{
	case true :
		fmt.Println("Nama Terlalu Panjang")

	case false :
		fmt.Println("Nama Sudah Benar")
	}


	name = "khannedykhannedykhannedy"

	length := len(name)
	switch {
	case length > 10 :
		fmt.Println("nama terlalu panjang")
	case length > 5 :
		fmt.Println("nama lumayan panjang")
	default :
		fmt.Println("nama sudah benar")
	}

}