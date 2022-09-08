package main

import (
	"fmt"
)

func PrimeNumber(nilai int) bool {
	var hasil bool
	hitung := 0
	for i := 1; i < nilai; i++ {
		if nilai%i == 0 {
			hitung++
		}
		if hitung >= 2 {
			hasil = false
		} else {
			hasil = true
		}
	}
	return hasil
}

func main() {
	fmt.Println(PrimeNumber(11))
	fmt.Println(PrimeNumber(13))
	fmt.Println(PrimeNumber(17))
	fmt.Println(PrimeNumber(20))
	fmt.Println(PrimeNumber(35))
	fmt.Println(PrimeNumber(36))
}
