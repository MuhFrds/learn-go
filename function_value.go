package main

import "fmt"

func getGoodBye(name string) string {
	return "Good bye " + name
}

func main() {

	contoh := getGoodBye
	misal := getGoodBye

	fmt.Println(contoh("Eko"))
	fmt.Println(misal("Eko"))
}