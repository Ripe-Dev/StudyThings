package main

import "fmt"

func main() {
	years := 15
	result := [3]int{0, 0, 0}

	result[0] = years

	switch years {
	case 1:
		result[1], result[2] = 15, 15

	case 2:
		result[1], result[2] = 24, 24
	}

	if years > 2 {
		result[1], result[2] = 24, 24
		for i := 3; i <= years; i++ {
			result[1] += 4
			result[2] += 5
		}
	}
	fmt.Println(result)

}
