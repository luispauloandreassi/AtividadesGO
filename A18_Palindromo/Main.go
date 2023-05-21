package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Bem vindo ao programa que verifica se o digitado é um Palindromo")
	fmt.Println("ATENCAO!!!! O programa faz diferenciacao etre Letras Maiusculas e Minusculas")
	fmt.Print("Informe uma palavra, frase ou número para serem verificados: ")

	scanner.Scan()
	dados := scanner.Text()

	fmt.Println("\nVerificando......")

	fmt.Println("TEXTO INFORMADO: " + dados)
	fmt.Println("TEXTO INVERTIDO: " + Invertedados(dados))
	VerificaPalindromo(dados, Invertedados(dados))

}
