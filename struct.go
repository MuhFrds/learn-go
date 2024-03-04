package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name )
}

func main() {
	var Firdaus Customer
	fmt.Println(Firdaus)

	Firdaus.Name = "Muhammad Firdaus"
	Firdaus.Address = "Indonesia"
	Firdaus.Age = 17
	fmt.Println(Firdaus)
	fmt.Println(Firdaus.Name)
	fmt.Println(Firdaus.Address)
	fmt.Println(Firdaus.Age)

	joko := Customer{
		Name: "joko",
		Address: "Indonesia",
		Age: 00,
	}

	fmt.Println(joko)

	budi := Customer{"Budi", "Indonesia", 00}
	fmt.Println(budi)

	budi.sayHello("Agus")
	Firdaus.sayHello("Agus")
	joko.sayHello("Agus")
}