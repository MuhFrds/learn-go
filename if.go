package main

import "fmt"

func main() {
	name := "Kurniawan"

	if name == "Eko" {
		fmt.Println("Hello Eko")
	} else if name == "Budi"{
		fmt.Println("Hello, Budi")
	} else if name == "Joko"{
		fmt.Println("Hello, Joko")
	}  else {
		fmt.Println("Hi, Boleh Kenalan?")
	}

	// length := len(name)
	// if	length > 5 {
	// 	fmt.Println("Nama Terlalu Panjang")
	// } else {
	// 	fmt.Println("Nama Sudahh Benar")
	// }

	// Short Statement
	if length := len(name);	length > 5 {
		fmt.Println("Nama Terlalu Panjang")
	} else {
		fmt.Println("Nama Sudahh Benar")
	}
}