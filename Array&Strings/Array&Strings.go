package main

import "fmt"

func ContarOcorrencias(palavras []string) map[string]int {

	contagem := make(map[string]int)

	for _, palavra := range palavras {
		contagem[palavra]++
	}

	return contagem
}

func main() {

	//	Array

	fmt.Println("Array & Slices")

	var array [3]string
	array[0] = "Zelda"
	array[1] = "Chan"
	array[2] = "Hey"

	fmt.Println(array[0], array[1])
	fmt.Println(array)
	fmt.Println(array[0:2])

	// Slice

	numPãesdeQueijo := make([]string, 10)
	// Qnt de Pãe de queijo e sua qualidade

	numPãesdeQueijo = []string{"Bonito", "Bonito", "Feio", "Vermelho", "Feio", "Feio"}

	numPãesdeQueijo = append(numPãesdeQueijo, "Feio", "Vermelho")

	numPãesdeQueijo = append(numPãesdeQueijo, "Vermelho", "Amarelo", "Crocante")

	//cabo :P
	resultado := ContarOcorrencias(numPãesdeQueijo)

	for palavra, quantidade := range resultado {
		fmt.Printf("%s: %d\n", palavra, quantidade)
	}

	fmt.Println(len(numPãesdeQueijo))
	fmt.Println(numPãesdeQueijo[3])
	fmt.Println(numPãesdeQueijo[7])

}

//	 Array & Slices
// Se trata basicamente de listas formadas por uma série de dados, no qual precisam
// 				obrigatóriamente pertencer ao mesmo tipo de retorno

//	Indices em Listas
// No caso do array e do slices, os indices das listas (isso é, a contagem de elementos)
// se conta o 0 como um elemento de sua lista
// ou seja em uma lista Array de 5 elementos, o primeiro pode ser puxado com o a[0]
// o segundo com o a[1], iadaiadaiada.

// Escrita facilitada
// "*nome da variael* := [4]int{2, 3, 4, 5}

//	Array
// Tamanho Fixo - o array obrigatóriamene precisa de uma prévia do tamanho da sua lista

//	Slices
// Tamanho Indefinido - diferente do array, a principal caracteristica de um slice é que não necessita de um
// numero limite na quantidade de elementos na lista, ainda assim, é possivel  limitar a lista com o comando
// "make"
