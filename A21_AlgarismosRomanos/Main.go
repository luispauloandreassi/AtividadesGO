package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Bem vindo ao programa que converte algarismos romanos.")
	fmt.Print("Informe um Algorismo Romano: ")
	scanner.Scan()
	algarismoRomano := scanner.Text()
	algarismoRomano = strings.ToUpper(algarismoRomano)

	if validarAlgarismosRomanos(algarismoRomano) {
		fmt.Println("\nALGARISMO ROMANDO        : " + algarismoRomano)
		fmt.Print("INTERIO CORRESPONDENTE   : ")
		fmt.Println(ConverteRomano(algarismoRomano))
	} else {
		fmt.Println("\nAlgarismo Informado é Inválido!!!")
		fmt.Println("O programa esta sendo finalizado...")
	}

}
