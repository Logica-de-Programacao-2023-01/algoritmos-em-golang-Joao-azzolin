package main

import "fmt"

func main() {
	var primeiro_número int
	var segundo_número int
	fmt.Print("Insira o primeiro número:")
	fmt.Scan(&primeiro_número)
	fmt.Print("Insira o segundo número:")
	fmt.Scan(&segundo_número)
	if primeiro_número > segundo_número {
		fmt.Println("x é maior que y")
	} else {
		fmt.Println("y é maior que x")
	}
}
