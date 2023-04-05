package main

import "fmt"

func main() {
	var base float64
	var altura float64
	var profundidade float64
	fmt.Print("Informe a base:")
	fmt.Scan(&base)
	fmt.Print("Informe a altura:")
	fmt.Scan(&altura)
	fmt.Print("Informe a profundidade:")
	fmt.Scan(&profundidade)
	fmt.Println("O volume da caixa é:", base*altura*profundidade)
}
