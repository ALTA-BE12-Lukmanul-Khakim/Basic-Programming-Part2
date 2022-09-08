package main

import "fmt"

func FaktorBilanganDesc(nilai int) string {
	var res string
	for i := nilai; i > 0; i-- {
		if nilai%i == 0 {
			res += fmt.Sprintln(i)
		}
	}
	return res
}

func main() {
	var nilai int
	fmt.Scanf("%d", &nilai)
	fmt.Println(FaktorBilanganDesc(nilai))
}
