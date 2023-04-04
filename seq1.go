package main

import "fmt"

func main() {
	var x int
	var y int
	var z int
	fmt.Print("Diga um número:")
	fmt.Scan(&x)
	fmt.Print("Diga outro número:")
	fmt.Scan(&y)
	fmt.Print("Diga mais um número:")
	fmt.Scan(&z)
	fmt.Println("A soma de x, y e z é:", x+y+z)
}
