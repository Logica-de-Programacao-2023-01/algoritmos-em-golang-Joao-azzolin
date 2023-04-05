package main

import "fmt"

func main() {
	var nome string
	var idade int
	var peso float64
	fmt.Print("Diga-me seu nome:")
	fmt.Scan(&nome)
	fmt.Print("Informe sua idade:")
	fmt.Scan(&idade)
	fmt.Print("Insira seu peso:")
	fmt.Scan(&peso)
	fmt.Println("Seu nome é:", nome, "Sua idade é:", idade, "Seu peso é:", peso)

}
