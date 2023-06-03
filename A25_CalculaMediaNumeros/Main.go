package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	continuar := true
	arrayNumeros := []float64{}

	for continuar == true {
		fmt.Println("Informe um número: ")
		scanner.Scan()
		numero, _ := strconv.ParseFloat(scanner.Text(), 64)

		arrayNumeros = append(arrayNumeros, numero)

		fmt.Println("Deseja continuar informando números: ")
		fmt.Println("1 - Sim.")
		fmt.Println("2 - Não.")
		scanner.Scan()
		continuar = scanner.Text() == "1"
	}

	var somaNumeros float64
	for _, num := range arrayNumeros {
		somaNumeros += float64(num)
	}

	resultado := somaNumeros / float64(len(arrayNumeros))
	fmt.Printf("A média dos números é: %.2f", resultado)
}
