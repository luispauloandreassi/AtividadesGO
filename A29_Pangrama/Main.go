package main

import (
	"bufio"
	"fmt"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	fmt.Println("Informe uma frase: ")
	scanner.Scan()
	frase := scanner.Text()

	var arrayLetras = make(map[string]string)
	for i := 0; i < len(frase); i++ {
		letra := string(frase[i])
		if _, existeLetra := arrayLetras[letra]; !existeLetra {
			arrayLetras[letra] = letra
		}
	}

	if len(arrayLetras) == 26 {
		fmt.Println("A frase informada é um pangrama")
	} else {
		fmt.Println("A frase informada não é um pangrama")
	}
}
