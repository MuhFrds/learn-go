package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1 // pointer

	address2.City = "Bandung"
	fmt.Println(address1) // ikut berubah
	fmt.Println(address2) // berubah menjadi Bandung

	// address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)

}