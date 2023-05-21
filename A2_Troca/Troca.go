package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Bem-vindo ao programa Troca Troca")
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Informe um valor: ")
	scanner.Scan()
	valorA := scanner.Text()

	fmt.Print("Informe outro valor: ")
	scanner.Scan()
	valorB := scanner.Text()

	fmt.Println("\n_____Valores Originais:")
	fmt.Println("Valor1:", valorA)
	fmt.Println("Valor2:", valorB)

	transitoria := valorA
	valorA = valorB
	valorB = transitoria

	fmt.Println("\n_____Valores Após a Troca:")
	fmt.Println("Valor1:", valorA)
	fmt.Println("Valor2:", valorB)
}
