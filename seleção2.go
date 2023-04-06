package main

import "fmt"

func main() {
	var x int
	var y int
	var z int
	fmt.Print("Digite o primeiro número x:")
	fmt.Scan(&x)

	fmt.Print("Digite o segundo número y:")
	fmt.Scan(&y)

	fmt.Print("Digite o terceiro número z:")
	fmt.Scan(&z)

	if x < y && x < z {
		fmt.Println("x é o menor número")

	} else if y < x && y < z {
		fmt.Println("y é o menor número")

	} else if z < x && z < y {
		fmt.Println("z pe o menor número")
	}
}
