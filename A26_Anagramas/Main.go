package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func sortString(str string) string {
	// Converter a string em uma lista de caracteres
	chars := strings.Split(str, "")

	// Ordenar a lista de caracteres
	sort.Strings(chars)

	// Juntar os caracteres ordenados em uma string
	return strings.Join(chars, "")
}

func main() {
	var primeiraPalavra = ""
	var segundaPalavra = ""

	fmt.Print("Informe a primeira palavra: ")
	scanner.Scan()
	primeiraPalavra = scanner.Text()

	fmt.Print("Informe a segunda palavra: ")
	scanner.Scan()
	segundaPalavra = scanner.Text()

	primeiraPalavra = strings.ToLower(primeiraPalavra)
	segundaPalavra = strings.ToLower(segundaPalavra)

	primeiraPalavraOrdenada := sortString(primeiraPalavra)
	segundaPalavraOrdenada := sortString(segundaPalavra)

	if primeiraPalavraOrdenada == segundaPalavraOrdenada {
		fmt.Println("As palavras informadas são anagramas.")
	} else {
		fmt.Println("As palavras informadas não são anagramas.")
	}
}
