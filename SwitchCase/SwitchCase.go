package main

import (
	"fmt"
)

// lipe, voce eh buro, mt mai facil assim kk
// // func pelucia(
// // 	QuntMembros: int
// // 	Cor : string
// // 	Quadrupede: bool
// // 	Pescoço: bool
// )

//	Switch Case
// pelo amor de deus, nunca trablhe com subjetividade, ou tu vai recodar uma pa de vez

func main() {

	QuntMembros := 4
	Cor := "Rosa"
	Quadrupede := false
	Pescoço := false

	Provavel := 0

	if Quadrupede == true {
		switch QuntMembros {
		case 3:
			fmt.Printf("É o Gato Branco\n")

		case 6:
			fmt.Printf("É o Poro\n")

		case 7:
			Provavel = 2
			fmt.Printf("Então ele tem %d membros e é Quadrupede\n", QuntMembros)
		}
		// Se não for quadrupede
	} else {
		if QuntMembros >= 4 {
			switch QuntMembros {
			case 4:
				Provavel = 1
				fmt.Printf("Então ele tem %d membros e n é Quadrupede\n", QuntMembros)

			case 7:
				Provavel = 3
				fmt.Printf("Então ele tem %d membros e n é Quadrupede\n", QuntMembros)
			}

		}
	}

	// Considerar as Provaveis

	if Provavel == 0 {
		fmt.Printf(":D")
	} else {
		switch Provavel {

		case 2:
			switch Cor {
			case "Cinza":
				fmt.Printf(" Descobri!\n Ele é o Gato Aborba\n :D")

			case "Rosa":
				fmt.Printf(" Descobri!\n Ele é o Gato Rosa da Sofi\n :D")

			}

		case 1:
			switch Cor {
			case "Creme":
				fmt.Printf(" Descobri!\n Ele é o Donout\n :D")

			case "Rosa":
				fmt.Printf(" Descobri!\n Ele é o Kirby\n :D")
			}
		case 3:
			if Pescoço == true {
				fmt.Printf(" Descobri!\n Ele é o Ganso\n :D")
			} else {
				fmt.Printf(" Descobri!\n Ele é o Cinnamon Roll\n :D")
			}
		}
	}
}
