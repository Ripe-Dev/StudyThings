package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, -11, -12, -13, -14, -15}
	var more int
	var less int
	for _, allnums := range numbers {
		if allnums < 0 {
			less += allnums
		} else if allnums > 0 {
			more += 1
		}
	}
	list := []int{more, less}
	fmt.Println(list)
}
