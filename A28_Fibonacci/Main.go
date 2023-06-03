package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	sequenciaNumeros := []int{}

	fmt.Println("Informe um número: ")
	scanner.Scan()
	numeroInformado, _ := strconv.Atoi(scanner.Text())

	numeroAntecessor, numeroAtual := 0, 1
	for numeroAtual <= numeroInformado {
		sequenciaNumeros = append(sequenciaNumeros, numeroAtual)
		numeroAntecessor, numeroAtual = numeroAtual, numeroAntecessor+numeroAtual
	}

	fmt.Println("Sequência: ", sequenciaNumeros)
}
