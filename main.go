package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("")
	fmt.Println("-----------------------------PROGRESSÃO GEOMÉTRICA-----------------------------------")

	var primeiroTermo float64
	fmt.Print("Digite o primeiro termo: ")
	fmt.Scan(&primeiroTermo)

	var razao float64
	fmt.Print("Digite a razao: ")
	fmt.Scan(&razao)

	var numeroDeElementos float64
	fmt.Print("Digite o número de elementos: ")
	fmt.Scan(&numeroDeElementos)

	for numeroDeElementos > 0 {
		numeroDeElementos -= 1.

		sequencia := primeiroTermo*(math.Pow(razao, numeroDeElementos))

		defer fmt.Print(sequencia, " ")
	}   

	
}