package main

import "fmt"

// somar o numero de vezes que a entrada pedir
// essa soma precia ser um numero por vez até o numero de entrada

func main() {

	num := 2
	numb := 0
	plus := 1
	for i := 0; i < num; i += 1 {
		numb += plus
		plus += 1
		fmt.Println(numb)
		fmt.Println(plus)
	}

	fmt.Printf("Conclusão: %d", numb)
}
