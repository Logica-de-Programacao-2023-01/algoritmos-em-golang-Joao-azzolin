package main

import "fmt"

func main() {
	var x int
	fmt.Print("Digite um número inteiro:")
	fmt.Scan(&x)
	if x%3 == 0 && x&5 == 0 {
		fmt.Println("Esse número é múltiplo de 3 e 5 simultaneamente")
	} else {
		fmt.Println("Esse número não é múltiplo de 3 e 5 simultaneamente")
	}
}
