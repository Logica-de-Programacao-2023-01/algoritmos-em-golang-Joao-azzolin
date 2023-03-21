package main

import "fmt"

func main() {
	var x int
	var y int
	var z int
	fmt.Print("Diga um número x:")
	fmt.Scan(&x)
	fmt.Print("Diga número y:")
	fmt.Scan(&y)
	fmt.Print("Diga um número z:")
	fmt.Scan(&z)
	if x < y && x < z {
		fmt.Println("x é o menor entre os 3 números")
	}
	if y < x && y < z {
		fmt.Println("y é o menor entre os 3 números")
	}
	if z < x && z < y {
		fmt.Println("z é o menor entre os 3 números")
	}

}
