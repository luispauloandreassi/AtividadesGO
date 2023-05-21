package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Bem-vindo ao programa FizzBuzz")

	fmt.Println("Favor informar um intervalo de números:")
	fmt.Print("Número Inicial: ")
	scanner.Scan()
	numeroInicialStr := scanner.Text()
	numeroInicial, _ := strconv.Atoi(numeroInicialStr)

	fmt.Print("Número Final: ")
	scanner.Scan()
	numeroFinalStr := scanner.Text()
	numeroFinal, _ := strconv.Atoi(numeroFinalStr)

	if numeroFinal <= numeroInicial {
		fmt.Println("ERRO!!! Numero Final é menor ou Igual ao Numero Inicial.")
		fmt.Println("O programa está sendo finalizado!!!!")
	} else {
		FizzBuzz(numeroInicial, numeroFinal)
	}

}
