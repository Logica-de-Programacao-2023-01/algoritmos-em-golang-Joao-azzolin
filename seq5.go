package main

import "fmt"

func main() {
	var idade float64
	fmt.Print("Diga sua idade:")
	fmt.Scan(&idade)
	fmt.Println("Sua idade em dias é:", idade*365)
}
