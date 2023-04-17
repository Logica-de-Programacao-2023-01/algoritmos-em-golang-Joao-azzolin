package main

import "fmt"

func main() {
	var x int
	fmt.Print("Digite um número inteiro:")
	fmt.Scan(&x)
	fmt.Println("Os divisores de", x, "são:")
	for i := 1; i < (x + 1); x++ {
		if x%i == 0 {
			fmt.Println(i)
		}
	}
}
