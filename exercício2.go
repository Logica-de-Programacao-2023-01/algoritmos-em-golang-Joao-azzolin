package main

import "fmt"

func main() {
	var base float64
	var altura float64
	fmt.Print("Insira a base é:")
	fmt.Scan(&base)
	fmt.Print("Insira a altura é:")
	fmt.Scan(&altura)
	fmt.Println("A área do retângulo será:", (base*altura)/2)
}
