package main

import (
	"fmt"
	"strconv"
	"strings"
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

func FullPrima(n int) bool {
	var (
		res  bool = true
		res2 bool = true
		set  string
	)

	res2 = PrimeNumber(n)

	if res2 {
		set = strconv.Itoa(n)         //konversi angka ke string
		sli := strings.Split(set, "") //proses split string

		for i := 0; i < len(sli); i++ {
			angka, _ := strconv.Atoi(sli[i]) //pengembalian angka

			if !PrimeNumber(angka) {
				return false
			}

		}
	}
	return res
}

func main() {
	fmt.Println(FullPrima(2))  // true
	fmt.Println(FullPrima(3))  // true
	fmt.Println(FullPrima(11)) // false
	fmt.Println(FullPrima(13)) // false
	fmt.Println(FullPrima(23)) // true
	fmt.Println(FullPrima(29)) // false
	fmt.Println(FullPrima(37)) // true
	fmt.Println(FullPrima(41)) // false
	fmt.Println(FullPrima(43)) // false
	fmt.Println(FullPrima(53)) // true
}
