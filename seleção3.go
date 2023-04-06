package main

import "fmt"

func main() {
	var número int
	fmt.Print("Digite um número:")
	fmt.Scan(&número)
	if número%2 == 0 {
		fmt.Println("Esse número é par")
	} else if número%2 != 0 {
		fmt.Println("Esse número é ímpar")
	}
}
