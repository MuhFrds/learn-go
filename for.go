package main

import "fmt"

func main() {

	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("perulanagan ke", counter)
	// 	counter++
	// }

	// for dengan init dan post statement
	// init statement di eksekusi pertama kali di awal
	// post statement diulang ulang di akhir perulangan 

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("perulanagan ke", counter)
	}

	fmt.Println("selesai")

	names := []string{"muh", "Muhammad", "Firdaus"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}

	for _, name := range names {
		fmt.Println(name)
	}
}