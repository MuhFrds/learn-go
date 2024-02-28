package main

import "fmt"

func main() {
	const (
		firstName string = "Muhammad"
		lastName         = "Firdaus"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	// const firstName string = "Muhammad"
	// const lastName = "Firdaus"

	//eror
	// firstName = "Budi"
	// lastName = "Joko"
}