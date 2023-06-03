package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Bem vindo ao programa TROCO: ")
	fmt.Print("Informe o valor da compra: ")
	scanner.Scan()
	valorCompraStr := scanner.Text()
	valorCompra, _ := strconv.ParseFloat(valorCompraStr, 64)

	fmt.Print("Informe o valor Pago: ")
	scanner.Scan()
	valorPagoStr := scanner.Text()
	valorPago, _ := strconv.ParseFloat(valorPagoStr, 64)

	//Calculando o valor do Troco
	valor := valorPago - valorCompra

	if valor < 0 {
		fmt.Println("\n-------------------")
		fmt.Println("      Valor Pago menor que a quantidade da compra...")
		fmt.Println("      Você deveria ter dado mais dinheiro!")
		fmt.Println("      O programa está sendo finalizado......")
		os.Exit(1)
	}

	//Chamando a função para retornar a quantidade menor de cedulas/moeda
	troco := CalcularCelulasMoedasTroco(valor)

	fmt.Printf(" \nTroco Calculado: R$%.2f:\n", valor)
	fmt.Println("Para esse troco, a menor quandidade de Cédulas/Moedas é: ")
	for valor, qtd := range troco {
		if valor < 1.00 {
			fmt.Printf(" R$%.2f Qtd Moeda(s)-> %d\n", valor, qtd)
		} else {
			fmt.Printf(" R$%.2f Qtd Cédula(s)-> %d\n", valor, qtd)
		}
	}
}

var valores = []float64{100, 50, 10, 5, 1, 0.5, 0.1, 0.05, 0.01}

func CalcularCelulasMoedasTroco(valor float64) map[float64]int {
	troco := make(map[float64]int)

	for _, vlr := range valores {
		qtd := int(valor / vlr)
		if qtd > 0 {
			troco[vlr] = qtd
			valor -= float64(qtd) * vlr
		}
	}
	return troco
}
