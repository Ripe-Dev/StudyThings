package main

import (
	"fmt"
	// "time"
)

func main() {

	// 	//Ex de looping simples
	// (Lembrete: se for usar uma var dentro dele, declarar fora do looping plz)

	sum := 0
	// esse ++ siginifica que sera adicionado mais 1 ao indice, para seguir para a proxima
	// posição, ou atividade
	for i := 0; i < 10; i++ {
		// esse += é igual a: sum = sum + i (o que faz sum -= i; sum = sum - i)
		sum += i
		fmt.Println(sum)
	}

	// // Ex de looping infinito travado por segundos
	// // (lembrete: Ctrl + c no terminal cancela o codigo)

	// for {
	// 	fmt.Println("Oi estou vivo :D")
	// 	time.Sleep(1 * time.Second)
	// }

	// For Range

	frutas := []string{"Abóbora", "Abacaxi", "Maçã", "Banana"}
	for _, fruta := range frutas {
		fmt.Println(fruta)
	}
}
