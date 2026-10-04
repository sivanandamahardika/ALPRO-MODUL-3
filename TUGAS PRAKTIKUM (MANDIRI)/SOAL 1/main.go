package main

import "fmt"

func main() {
	var jumlah_koin, emas, perak, tembaga, sisa int

	fmt.Scan(&jumlah_koin)

	emas = jumlah_koin / 9
	sisa = jumlah_koin % 9

	perak = sisa / 3
	tembaga = sisa % 3

	fmt.Println(emas, perak, tembaga)
}