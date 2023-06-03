package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	fmt.Println("Informe a cotação do dólar: ")
	scanner.Scan()
	cotacao, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Println("Informe o saldo disponível: ")
	scanner.Scan()
	saldoDisponivel, _ := strconv.ParseFloat(scanner.Text(), 64)

	valorConvertido := saldoDisponivel * cotacao

	fmt.Println("Valor convertido: ", valorConvertido)
}
