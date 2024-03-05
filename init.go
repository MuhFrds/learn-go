package main

import (
	"fmt"

	"helloworld.go/database"
	_ "helloworld.go/internal"
)

func main(){
	fmt.Println(database.GetDatabase())
}