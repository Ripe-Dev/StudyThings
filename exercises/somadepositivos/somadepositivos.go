package main

import "fmt"

func main() {

	numbers := []int{1, -4, 7, 12, -2, 3}
	plus := 0
	for _, allnums := range numbers {
		if allnums > 0 {
			plus += allnums
			fmt.Println(plus)
		}
	}
	fmt.Println(plus)
}
