package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	fmt.Println("Bem-vindo ao programa que calcula o Ajuste Salarial")
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Informe o Valor de um Salario: ")
	scanner.Scan()
	salarioStr := scanner.Text()
	salario, _ := strconv.ParseFloat(salarioStr, 64)

	fmt.Print("Informe o percentual de reajuste salarial: ")
	scanner.Scan()
	percentualStr := scanner.Text()
	percentual, _ := strconv.ParseFloat(percentualStr, 64)

	fmt.Println("\nNOVO SALARIO:", CalcularSalario(salario, percentual))

}
