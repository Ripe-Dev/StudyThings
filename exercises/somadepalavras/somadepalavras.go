package main

import "fmt"

func main() {
	n := 0
	s := "ola"
	sb := s
	for i := 1; i < n; i++ {
		s += sb

	}
	fmt.Println(s)

}
