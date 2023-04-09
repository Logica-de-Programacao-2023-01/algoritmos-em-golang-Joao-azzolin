package main

import "fmt"

func main() {
	var x int
	fmt.Print("Escolha um número:")
	fmt.Scan(&x)
	fmt.Println(x, "É um ótimo número para operarmos!")
	for i := 1; i >= 1; i++ {
		if x%i == 0 {
			fmt.Println(i)
		}
	}
}
