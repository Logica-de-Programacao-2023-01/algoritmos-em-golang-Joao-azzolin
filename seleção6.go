package main

import "fmt"

func main() {
	var x int
	var y int
	fmt.Print("Insira um número:")
	fmt.Scan(&x)
	fmt.Print("Insira outro número:")
	fmt.Scan(&y)
	if x > 0 && y > 0 {
		fmt.Println("O resultado da multiplicação é:", x*y)
	} else if x < 0 || y < 0 {
		fmt.Println("O resultado da soma é:", x+y)
	}
}
