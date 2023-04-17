package main

import "fmt"

func main() {
	var x int
	fmt.Print("Digite um número:")
	fmt.Scan(&x)
	for i := 1; i <= 10; i++ {
		fmt.Println(x * i)
	}

}
