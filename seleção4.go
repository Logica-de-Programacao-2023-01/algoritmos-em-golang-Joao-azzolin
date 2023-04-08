package main

import "fmt"

func main() {
	var altura float64
	var peso float64
	var sexo string

	fmt.Print("Digite sua altura em metros:")
	fmt.Scan(&altura)
	fmt.Print("Digite seu peso em kgs:")
	fmt.Scan(&peso)
	fmt.Print("Digite seu sexo(M para masculino && F para feminino:")
	fmt.Scan(&sexo)

	if sexo == "M" {
		fmt.Println("Peso ideal ==", (0.75*altura)-62.5)
	} else if sexo == "F" {
		fmt.Println("Peso ideal ==", (0.67*altura)-56)
	}
	IMC := peso / (altura * altura)
	if IMC < 18.5 {
		fmt.Println("Você está abaixo do peso ideal")
	} else if IMC > 25 {
		fmt.Println("Você está acima do peso ideal")
	} else {
		fmt.Println("Você esta dentro do peso ideal")
	}
}
