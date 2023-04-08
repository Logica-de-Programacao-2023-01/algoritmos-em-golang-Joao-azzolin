package main

import "fmt"

func main() {
	var x float64
	var y float64
	var z float64
	fmt.Print("Digite o primeiro número:")
	fmt.Scan(&x)
	fmt.Print("Digite o segundo número:")
	fmt.Scan(&y)
	fmt.Print("Digite o terceiro número:")
	fmt.Scan(&z)
	if x > y && x > z && y > z {
		fmt.Println("A  leitura dos três números em ordem crescente é:", x, y, z)
	} else if x > y && x > z && z > y {
		fmt.Println(" leitura dos três números em ordem crescente é:", x, z, y)
	} else if y > z && y > x && x > z {
		fmt.Println(" A leitura dos três números em ordem crescente é:", y, x, z)
	} else if y > z && y > x && z > x {
		fmt.Println("A leitura dos três números em ordem crescente é:", y, z, x)
	} else if z > x && z > y && y > x {
		fmt.Println("A  leitura dos três números em ordem crescente é:", z, y, x)
	} else if z > x && z > y && x > y {
		fmt.Println("A leitura dos três números em ordem crescente é:", z, x, y)
	}

}
