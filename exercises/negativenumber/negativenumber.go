package main

import "fmt"

func main() {

	// transformar o algoritimo positivo em negativo
	// sem adulterar o negativo

	num := 0

	if num > 0 {

		multiply := num * 2

		num -= multiply

		fmt.Println(num)
	} else {

		fmt.Printf("O numero %d ja é negativo", num)
	}

	// fmt.Println(num)
}
