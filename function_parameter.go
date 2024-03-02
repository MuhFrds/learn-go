package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	sayHelloTo("Muhammad", "Firdaus")
	sayHelloTo("Eko", "Khannedy")
}