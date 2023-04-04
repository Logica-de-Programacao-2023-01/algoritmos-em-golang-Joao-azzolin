package main

import "fmt"

func main() {
	var x int
	var y int
	fmt.Print("Digite a quantidade de dias que o funcionário trabalha:")
	fmt.Scan(&x)
	fmt.Print("Digite o valor da diária do funcionário:")
	fmt.Scan(&y)
	fmt.Println("Seu salário será igual à:", x*y)
}
