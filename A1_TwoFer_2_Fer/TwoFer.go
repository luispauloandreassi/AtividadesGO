package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Bem-vindo ao programa TwoFer")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Informe um nome qualquer: ")
	scanner.Scan()
	nome := scanner.Text()
	twofer(nome)
}
