package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Bem vindo ao programa que calcula fatorial: ")
	fmt.Print("Informe um Número quaquer: ")

	scanner.Scan()
	numeroStr := scanner.Text()
	numero, _ := strconv.ParseFloat(numeroStr, 64)

	CalculaFatorial(numero)

}
