package main

import "fmt"

func main() {
	var x float64
	fmt.Print("Insira seu peso em kg:")
	fmt.Scan(&x)
	fmt.Println("Seu peso em libras é:", x*2.2046)
}
