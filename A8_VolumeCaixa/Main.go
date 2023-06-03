package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	fmt.Println("Informe o comprimento: ")
	scanner.Scan()
	comprimento, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Println("Informe a largura: ")
	scanner.Scan()
	largura, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Println("Informe a altura: ")
	scanner.Scan()
	altura, _ := strconv.ParseFloat(scanner.Text(), 64)

	volume := comprimento * largura * altura
	fmt.Println("O volume da caixa é: ", volume)
}
