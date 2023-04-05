package main

import "fmt"

func main() {
	var x float64
	var y float64
	var z float64
	var h float64
	fmt.Print("Insira o primeiro número:")
	fmt.Scan(&x)
	fmt.Print("Insira o segundo número:")
	fmt.Scan(&y)
	fmt.Print("Insira o terceiro número:")
	fmt.Scan(&z)
	fmt.Print("Insira o quarto número:")
	fmt.Scan(&h)
	fmt.Println("A média aritmética é:", (x+y+z+h)/4)
}
