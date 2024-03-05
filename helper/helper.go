package helper

import "fmt"

// nama variable atau function yang diawali huruf kecil tidak bisa diakses diluar package
var version = "1.0.0"
var Appliction = "golang"

func sayGoodBye(name string) string {
	return "Good Bye " + name
}

func SayHello(name string) string {
	return "Hello " + name
}

func Contoh(){
	sayGoodBye("Eko")
	fmt.Println(version)
}