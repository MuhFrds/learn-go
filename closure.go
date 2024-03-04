package main

import "fmt"

func main() {


	// harap menggunakan fitur closure dengan bijak	

	counter := 0

	
	increment := func() {
		fmt.Println("increment")
		counter ++
	}


	increment()
	increment()
	increment()

	fmt.Println(counter)
}