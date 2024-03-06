package main

import (
	"errors"
	"fmt"
)
func pembagian(nilai int, pembagi int) (int, error){
	if pembagi == 0 {
		return 0, errors.New("Pembagian dengan nol")
	}else {
		return nilai / pembagi , nil
	}
}

func main (){
	hasil, err := pembagian(100, 0)
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else{
		fmt.Println("Error", err.Error())
	}

}