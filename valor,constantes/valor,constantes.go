package main

import "fmt"

func main() {
	// Variaveis
	// var + nome da variavel + tipo
	var nome string
	nome = "bento"
	fmt.Println(nome)

	nome = "lipe"
	fmt.Println(nome)

	var a = 1
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b)
	fmt.Println(c)

	a = 2
	fmt.Println(a)

	oi := "pear"
	fmt.Println(oi)

	// Constante (Não auteravel pós declaração)

	const idadeBento = 4
	fmt.Println(idadeBento)

}
