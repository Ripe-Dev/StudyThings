package main

import "fmt"

//	Structs
// Forma de criar sua própria estrutura de dados
// Personalizar de acordo com a necesidade
// Podemos usar vários tipos diferentes

// "type *nome da estrutura* struct { *campos* }"
type Pelucia struct {
	Nome       string
	Cor        string
	QntMembros int
	Quadrupede bool
}

func main() {

	// fmt.Println(Pelucia{"Poro", "Branco", 6, true})
	// fmt.Println(Pelucia{Nome: "Kirby", Cor: "Rosa", QntMembros: 4, Quadrupede: false})
	// fmt.Println(Pelucia{Nome: "CinamonRoll", Quadrupede: false})
	// fmt.Println(Pelucia{QntMembros: 5})

	Num1 := Pelucia{Nome: "Gato+Abórba", Cor: "Cinza,Laranja", QntMembros: 5, Quadrupede: true}
	fmt.Println(Num1.Nome)

	Num2 := Pelucia{Nome: "Docinho", Cor: "Rosa", QntMembros: 5, Quadrupede: true}
	Num1.Cor = "Cinza, Laranja e Verde"

	Pelucias := []Pelucia{}
	Pelucias = append(Pelucias, Num1, Num2)
	fmt.Println(Pelucias)

	Participantes := map[int][]Pelucia{}
	Participantes[1] = Pelucias

	fmt.Println(Participantes)

}

// participantes := map[int]string{

// 	1: "Zelda",
// 	2: "Arya",
// 	3: "Kirby",
// 	4: "Poro",
// 	5: "Gato + Abóbora",
// 	6: "CinamonRoll",
// }
