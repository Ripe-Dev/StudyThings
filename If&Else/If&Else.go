package main

import (
	"fmt"
)

// IF / ELSE

func main() {

	pelucias := map[string]int{
		"Zelda":    1,
		"Cinnamon": 2,
		"Gato":     3,
	}

	pelucias["Kirby"] = 4
	pelucias["Poro"] = 5

	// if (condição) {	<ação>	}
	if len(pelucias) == 5 { // true
		fmt.Println("Ta chei")

	} else {

		valoratual := len(pelucias)
		fmt.Println("Põe mais", 5-valoratual)

	}

	//	Aritimética no go
	// Soma: 1 + 1
	// Subtação: 2-1
	// Multiplicação: 2 * 2
	// Divião: 10 / 2
	// Resto de Divisão: 10 % 2

	fmt.Println(11 % 3)
	fmt.Println(11 / 3)

	oi := 10
	if oi%2 == 0 {
		fmt.Printf("%d é par\n", oi)
	} else {
		fmt.Printf("%d é impar\n", oi)
	}

	tchau := "bye"
	hey := "hi"
	if oi%2 == 0 {
		fmt.Printf("%s how's goin?\n", hey)
	} else {
		fmt.Printf("%s, see u later\n", tchau)
	}

}

//	Formatação pelo Printf
// Uma das maneiras de formatar o retorno do código, que pode se utilizar
// das variaveis, é atravez do Printf que consiste em certos caracteres
// seguidos de uma letra, e cada um tem um proposito.
//
// 	%d -> Ele escreve o valor de uma variavel que seja um numero
// 				Codigo - Printf{"%d", *valor*}
//
//	%s -> Ele escreve o valor de uma variavel que seja uma string
// 				Codigo - Printf{"%s", *valor*}
//
//	\n -> Ele pula a linha, dando um "enter" na saida do terminal
//				Código - Printf{"*conteudo* \n"}
//	(Importante. Precisa se for usar em string, precisa estar dentro
// 						  para Funcionar :D)
