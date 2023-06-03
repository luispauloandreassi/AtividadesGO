package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	fmt.Println("Informe o primeiro número do intervalo: ")
	scanner.Scan()
	primeiroNumero, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Informe o segundo número do intervalo: ")
	scanner.Scan()
	segundoNumero, _ := strconv.Atoi(scanner.Text())

	resultado := 0
	for i := primeiroNumero; i <= segundoNumero; i++ {
		if i%2 == 0 {
			resultado = resultado + i
		}
	}

	fmt.Println("Soma dos números pares: ", resultado)
}
