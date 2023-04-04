package main

import "fmt"

func main() {
	var x float64
	var y float64
	var z float64
	fmt.Print("Insira o primeiro número:")
	fmt.Scan(&x)
	fmt.Print("Insira o segundo número:")
	fmt.Scan(&y)
	fmt.Print("Insira o terceiro termo")
	fmt.Scan(&z)
	fmt.Println("A média ponderada é:", (2*x+3*y+5*y)/2+3+5)

}
