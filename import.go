package main

import (
	"fmt"

	"helloworld.go/helper"
)

func main() {
	result := helper.SayHello("Eko")
	fmt.Println(result)

	fmt.Println(helper.Appliction)
	// fmt.Println(helper.version) // tidak bisa diakses
	// fmt.Println(helper.sayGoodBye("Eko")) // tidak bisa diakses
} 