package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Bem Vindo ao programa que calcula o score de uma letra")
	fmt.Print("Informe uma palavra: ")
	scanner.Scan()
	palavra := scanner.Text()

	CalculoScore(palavra)

}
