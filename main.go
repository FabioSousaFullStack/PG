package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("")
	fmt.Println("-----------------------------CALCULADORA DE PROGRESSÃO GEOMÉTRICA-----------------------------------")

	var primeiroTermo float64
	fmt.Print("Digite o primeiro termo: ")
	fmt.Scan(&primeiroTermo)

	var razao float64
	fmt.Print("Digite a razao: ")
	fmt.Scan(&razao)

	var numeroDeElementos float64
	fmt.Print("Digite o número de elementos: ")
	fmt.Scan(&numeroDeElementos)

	var expoente float64

	sequencia := []float64{}

	for expoente < numeroDeElementos {
		expoente += 1

		TermoGeral := primeiroTermo*(math.Pow(razao, expoente))

		sequencia = append(sequencia, TermoGeral)
		
	}   

	fmt.Println("Sequência ", sequencia)

	
}

