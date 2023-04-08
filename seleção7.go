package main

import "fmt"

func main() {
	var x float64
	fmt.Print("Digite o salário do funcionário:")
	fmt.Scan(&x)
	if x < 1000 {
		fmt.Println("O salário do funcionário com um aumento de 10% será:", x+(x*10/100))
	} else if x > 1000 {
		fmt.Println("O salário do funcionário com um aumento de 5% será:", x+(x*5/100))
	}
}
