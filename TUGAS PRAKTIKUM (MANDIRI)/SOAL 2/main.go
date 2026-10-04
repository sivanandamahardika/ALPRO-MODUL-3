package main

import "fmt"

func main() {
	var Gp, lembur, Gb, Bl, Pt int

	fmt.Scan(&Gp, &lembur)

	Bl = lembur * 45000
	Pt = Gp * 55 / 1000
	Gb = Gp + Bl - Pt

	fmt.Println(Gb)
}