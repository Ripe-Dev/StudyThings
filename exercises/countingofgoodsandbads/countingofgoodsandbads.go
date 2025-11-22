package main

import "fmt"

// scissors W Paper, l Rock

func main() {
	x := []string{"good", "bad", "good", "good", "bad", "good", "bad", "bad", "good", "bad", "bad"}
	var g int
	for _, list := range x {
		if list == "good" {
			g += 1
		}
	}
	fmt.Println(g)

	if g == 1 || g == 2 {
		fmt.Println("Publish")
	}

	if g > 2 {
		fmt.Println("I smell a series!")
	}
}
