package main

import (
	"fmt"
)

func main() {
	// Bool (True/False)
	fmt.Printf("type: %T-value: %v\n", true, true)

	// String - sequencia de bytes (tudo que tiver entre aspas)

	fmt.Printf("Type: %T-value: %v\n", "oi123", "oi123")

	// Int - Numero inteio escrito sem aspas
	fmt.Printf("Type: %T-value: %v\n", 1, 1)

	// Float - um numero que possua casas decimais/virgula (float64/float32)
	fmt.Printf("(Type: %T-value: %v\n", 1.233, 1.233)
}
