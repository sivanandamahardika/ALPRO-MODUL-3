package main

import "fmt"

func main() {
	var N int
	var jam, menit, detik int

	fmt.Scan(&N)

	jam = N / 3600
	menit = (N % 3600) / 60
	detik = N % 60

	fmt.Println(jam, menit, detik)
}