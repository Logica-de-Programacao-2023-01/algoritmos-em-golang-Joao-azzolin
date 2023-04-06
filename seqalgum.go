package main

import "fmt"

func main() {
	var idade float64
	fmt.Print("Qual é a sua idade?")
	fmt.Scan(&idade)
	fmt.Println("A sua idade em dias é:", idade*365)

}
