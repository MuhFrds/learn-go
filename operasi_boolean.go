package main

import "fmt"

func main() {

	var nilaiAkhir = 90
	var absensi = 80

	var lulusNilaiAkhhir bool = nilaiAkhir > 80
	var lulusAbsensi bool = absensi > 80

	var lulus bool = lulusNilaiAkhhir && lulusAbsensi

	fmt.Println(lulus)
}