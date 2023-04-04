package main

import "fmt"

func main() {
	var salário float64
	fmt.Print("Informe seu salário:")
	fmt.Scan(&salário)
	fmt.Println("Seu salário com aumento será:", salário+(salário*15/100))
}
