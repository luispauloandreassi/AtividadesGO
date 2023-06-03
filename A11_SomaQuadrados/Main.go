package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	fmt.Println("Informe o primeiro número: ")
	scanner.Scan()
	primeiroNumero, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Println("Informe o segundo número: ")
	scanner.Scan()
	segundoNumero, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Println("Informe o terceiro número: ")
	scanner.Scan()
	terceiroNumero, _ := strconv.ParseFloat(scanner.Text(), 64)

	resultadoPrimeiroNumero := primeiroNumero * primeiroNumero
	resultadoSegundoNumero := segundoNumero * segundoNumero
	resultadoTerceiroNumero := terceiroNumero * terceiroNumero
	resultado := resultadoPrimeiroNumero + resultadoSegundoNumero + resultadoTerceiroNumero

	fmt.Println("A soma dos quadrados é: ", resultado)
}
