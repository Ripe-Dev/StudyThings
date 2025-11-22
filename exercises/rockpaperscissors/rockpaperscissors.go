package main

import "fmt"

// scissors W Paper, l Rock

func main() {
	var p1, p2 string

	p1 = "paper"

	p2 = "paper"

	switch p1 {
	case "scissors":
		switch p2 {
		case "rock":
			fmt.Println("Player 2 won!")
		case "paper":
			fmt.Println("Player 1 won!")
		case "scissors":
			fmt.Println("Draw")
		}
	case "rock":
		switch p2 {
		case "scissors":
			fmt.Println("Player 1 won!")
		case "paper":
			fmt.Println("Player 2 won!")
		case "rock":
			fmt.Println("Draw")
		}

	case "paper":
		switch p2 {
		case "rock":
			fmt.Println("Player 1 won!")
		case "scissors":
			fmt.Println("Player 2 won!")
		case "paper":
			fmt.Println("Draw")
		}
	}
}
