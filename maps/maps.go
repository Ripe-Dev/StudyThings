package main

import "fmt"

func main() {

	idade := map[string]int{}
	idade["Zelda"] = 3
	idade["Poro"] = 2
	fmt.Println(idade)
	fmt.Println(idade["Poro"])
	fmt.Println(idade["Zelda"])

	participantes := map[int]string{

		1: "Zelda",
		2: "Arya",
		3: "Kirby",
		4: "Poro",
		5: "Gato + Abóbora",
		6: "CinamonRoll",
	}

	fmt.Println(participantes[5])
	fmt.Println(participantes)

	// Pontuação importante, sempre ficar atento ao tipo da "key" e do "value", para
	// não ficar perdido por completo na hora que der algo de errado

	participantes[7] = "Donout"
	fmt.Println(participantes)

	//	Maps
	// Os maps são listas, porém diferentes em um quesito, são listas no qual
	// difeente do array e dos slices, eles podem conter mais de um tipo, e anexar
	// valores de um tipo ao outro, usando do sistema de "key - value"

}
