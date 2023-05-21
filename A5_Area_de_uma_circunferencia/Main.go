package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	pi := 3.14159265
	fmt.Println("Bem-vindo ao programa que calcula Circunferencia")

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Informe o valor do raio: ")
	scanner.Scan()
	raioStr := scanner.Text()
	raio, _ := strconv.ParseFloat(raioStr, 64)

	fmt.Println("Valor da Circunferencia:", CalcularCircunferencia(pi, raio))

}
