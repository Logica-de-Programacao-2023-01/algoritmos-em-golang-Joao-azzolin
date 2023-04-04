package main

import "fmt"

func main() {
	var x float64
	fmt.Print("Insira um número:")
	fmt.Scan(&x)
	fmt.Println("Seu antecessor é:", x-1, "Seu sucessor é:", x+1)

}
