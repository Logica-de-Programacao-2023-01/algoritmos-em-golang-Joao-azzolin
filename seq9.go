package main

import "fmt"

func main() {
	var x float64
	fmt.Print("Digite o preço do produto:")
	fmt.Scan(&x)
	fmt.Println("Seu valor com 10% de desconto será:", x+(x*10/100))

}
