package main

import "fmt"

func main() {
	var reais float64
	var dólar float64
	fmt.Print("O valor em reais é:")
	fmt.Scan(&reais)
	fmt.Println("O valor em reais convertido para dólares é:", reais/5.03)
	fmt.Print("O valor em dólares é:")
	fmt.Scan(&dólar)
	fmt.Println("O valor em dólares convertido para reais é:", dólar*5.03)

}
