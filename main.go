package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("")
	fmt.Println("-----------------------------CALCULADORA DE PROGRESSÃO GEOMÉTRICA-----------------------------------")
	fmt.Println("")
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
	var soma float64

	sequencia := []float64{}

	for expoente < numeroDeElementos {
		
		expoente += 1
		termoGeral := primeiroTermo*(math.Pow(razao, expoente - 1))
		sequencia = append(sequencia, termoGeral)
		soma += (sequencia[int(expoente - 1)])
		
	}   

	fmt.Println("")
	fmt.Println("---------------------------PG--------------------------------------------------")
	fmt.Println("")
	fmt.Println("PG: ", sequencia)
	
	fmt.Println("Soma: ", soma)
	fmt.Println("")



	fmt.Println("-------------------TERMO GERAL-------------------------------------------")
	fmt.Println("")
	fmt.Print("Digite a posição do elemento da pg que gostaria de consultar: ")
	fmt.Scan(&expoente)
	termoGeral := primeiroTermo*(math.Pow(razao, expoente - 1))
	
	fmt.Println("")
	fmt.Println("O elemento que ocupa a posição", expoente , " é o ", termoGeral)
	fmt.Println("")
	fmt.Println("------------------O SISTEMA AGRADECE ESPERO QUE TENHA APRENDIDO----------")
	fmt.Println("")

}

