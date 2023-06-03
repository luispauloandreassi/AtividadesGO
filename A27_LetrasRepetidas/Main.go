package main

import (
	"bufio"
	"fmt"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	fmt.Println("Informe uma palavra: ")
	scanner.Scan()
	palavra := scanner.Text()

	arrayLetras := make(map[byte]int)

	for i := 0; i < len(palavra); i++ {
		letra := palavra[i]
		arrayLetras[letra]++
	}

	var letraNaoRepetida byte
	var letraNaoRepetidaEncontrada bool

	for i := 0; i < len(palavra); i++ {
		letra := palavra[i]
		if arrayLetras[letra] == 1 {
			letraNaoRepetida = letra
			letraNaoRepetidaEncontrada = true
			break
		}
	}

	if letraNaoRepetidaEncontrada {
		fmt.Printf("A primeira letra não repetida na palavra é: %c\n", letraNaoRepetida)
	} else {
		fmt.Println("Não existe alguma letra que não se repete na palavra informada.")
	}
}
