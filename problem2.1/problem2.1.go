package main

import "fmt"

func FaktorBilangan(nilai int) string {
	var res string
	for i := 1; i <= nilai; i++ {
		if nilai%i == 0 {
			res += fmt.Sprintln(i)
		}
	}
	return res
}

func main() {
	var nilai int
	fmt.Scanf("%d", &nilai)
	fmt.Println(FaktorBilangan(nilai))
}
