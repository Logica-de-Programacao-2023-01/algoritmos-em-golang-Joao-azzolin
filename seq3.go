package main

import "fmt"

func main() {
	var peso float64
	var altura float64
	fmt.Print("Insira seu peso em kg:")
	fmt.Scan(&peso)
	fmt.Print("Insira sua altura em metros:")
	fmt.Scan(&altura)
	fmt.Println("Seu IMC é:", peso/(altura*altura))

}
