package main

import "fmt"

func main() {
	var x int
	fmt.Print("Insira um número qualquer:")
	fmt.Scan(&x)

	for i := 1; i < 11; i++ {
		fmt.Println(x * i)
	}
}
