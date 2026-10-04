package main

import "fmt"

func main() {
	var N, hari int

	fmt.Scan(&N)

	hari = (4 + N) % 7

	if hari == 0 {
		hari = 7
	}

	fmt.Println(hari)
}