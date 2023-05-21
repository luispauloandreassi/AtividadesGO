package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Bem vindo ao programa que calcula se um numero é primo: ")
	fmt.Println("I N F O R M A C A O:\n     Número Primo é aquele divisivel apenas por 1 e por ele mesmo.")
	fmt.Print("Informe um Número inteiro para ser verificado: ")

	scanner.Scan()
	numeroStr := scanner.Text()
	numero, _ := strconv.Atoi(numeroStr)

	if VerificaPrimo(numero) {
		fmt.Println("=> O numero ", numero, "É primo")
	} else {
		fmt.Println("=> O numero ", numero, "Não é primo")
	}

}
