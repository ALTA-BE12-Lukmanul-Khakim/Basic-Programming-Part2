package main

import "fmt"

func KonversiNilai(nilai int) string {
	var huruf string
	if nilai >= 80 && nilai <= 100 {
		huruf = ("A")
	} else if nilai >= 65 && nilai <= 79 {
		huruf = ("B+")
	} else if nilai >= 50 && nilai <= 64 {
		huruf = ("B")
	} else if nilai >= 35 && nilai <= 49 {
		huruf = ("C")
	} else {
		huruf = ("D")
	}
	return huruf
}

func main() {
	var nilai int = 70

	fmt.Println(KonversiNilai(nilai))
}
