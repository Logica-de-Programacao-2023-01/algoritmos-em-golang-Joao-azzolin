package main

import "fmt"

func main() {
	var x int
	fmt.Print("Digite a idade de uma pessoa:")
	fmt.Scan(&x)
	if x <= 9 {
		fmt.Println("Essa pessoa é mirim:")
	} else if x >= 10 && x <= 13 {
		fmt.Println("Essa pessoa é infantil:")
	} else if x >= 14 && x <= 17 {
		fmt.Println("Essa pessoa é juvenil:")
	} else if x >= 18 {
		fmt.Println("Essa pessoa é adulta:")
	}
}
