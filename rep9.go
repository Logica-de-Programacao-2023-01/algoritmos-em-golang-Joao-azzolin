package main

import "fmt"

func main() {
	var soma int
	var quantidade int
	for {
		var x int
		fmt.Print("Diga um número (0 para parar).")
		fmt.Scan(&x)
		if x == 0 {
			break
		}
		soma += x
		quantidade++
	}
	if quantidade > 0 {
		média := soma / quantidade
		fmt.Println("A média é:", média)
	}
}
