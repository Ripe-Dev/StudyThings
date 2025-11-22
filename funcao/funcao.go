package main

import "fmt"

// Finalmente agr eu tenho uma certeza mlhr do pra que servia esses
// parentesis depois do nome da função kk

// Basicamente, puxar uma função existente vai fazer o código ja saber
// o que precsa ser feito, caso contrario você porecisa estipular a
// (ironicamente) FUNÇÃO dessa função kk, você faz isso colocando

// "func *nome da fução* (*valor* *tipo do valo..r*, *outros PARAMETROS*) *tipo do reorno* {*codigo*}"

// Importante anotar que é extremamente necessario que você coloque os tipos
// nesses locais especificados, que é uma pratica boa e obrigatória

func soma(x int, y int) int {

	return x + y

}

func subtração(x int, y int) int {

	return x - y

}

func PrintaNome(nome string) (string, string) {

	return nome, nome

}

func PrintNomeCompleto(nome, sobrenome string) (string, string) {

	return nome, sobrenome

}

func main() {
	soa := soma(12, 21)
	fmt.Println(soa)

	soa = soma(31, 40)
	fmt.Println(soa)

	sub := subtração(10, 5)
	fmt.Println(sub)

	// ponto legal, caso eu queira não utilizar um dos dois valores
	// eu posso simplesmente colocar um underline "_" que da certo tambem
	nome1, _ := PrintaNome("Lipe")
	fmt.Println(nome1)

	nome, sobrenome := PrintNomeCompleto("Lipe", "Rodrigues")
	fmt.Println(sobrenome)
	fmt.Println(nome)

}

// Info legal, caso a função criada em um arquivo comece com sua
// letra inicial maiuscula, ela pode ser usada em outro pacoe, dando
// import no pacote do arquivo original (ou criando um arquivo
// nesse pacote com a função criada), basta colocar
// "*nome do pacote*.*nome da função*()"

// mas caso queira que a função não seja PÚBLICA, basta trocar o
// primeiro caractere de maiusculo para minusculo, tornando-la
// PRIVADA
