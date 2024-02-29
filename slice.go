package main

import (
	"fmt"
)

func main() {

	names := [...]string{"Eko", "Kurniawan", "Khannedy", "Joko", "Budi", "Nugraha"}

	slice1 := names[4:6]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[3:]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)

	days := [...]string {"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}

	daySlice1 := days[5:]
	fmt.Println(daySlice1)

	daySlice1[0] = "sabtu baru"
	daySlice1[1] = "minggu baru"
	fmt.Println(daySlice1)
	fmt.Println(days) //[senin, selasa, rabu, kamis, jumat, sabtu baru, minggu baru]

	// append menambahkan data baru ke array jika array penuh akan dibuat array baru secara otomatis
	daySlice2 := append(daySlice1, "Libur baru")
	daySlice2[0] = "sabtu lama"
	// days := [...]string {"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu", "Libur baru"}
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)


	// newSlice := make([]string, 2,5)
	var newSlice []string = make([]string, 2,5)

	newSlice[0] = "Eko"
	newSlice[1] = "Eko"
	// newSlice[2] = "Eko" // error, harusnya menggunakan append

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Eko")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))


	newSlice2[0] = "Budi"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// bedanya array dan slice adalah array ada deklarasi jumlah data sedangkan slice tidak ada
	iniArray := [...]int{1,2,3,4,5} 
	iniSlice := []int{1,2,3,4,5} 

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}