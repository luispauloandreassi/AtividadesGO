package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	fmt.Println("------- Bem-vindo ao LagartoSpock ;) -------\n")

	fmt.Println("Informe o nome do 1° jogador: ")
	scanner.Scan()
	primeiroJogador := scanner.Text()

	fmt.Println("Informe o nome do 2° jogador: ")
	scanner.Scan()
	segundoJogador := scanner.Text()

	fmt.Println(primeiroJogador, ", escolha uma opção: ")
	fmt.Println("1-Pedra")
	fmt.Println("2-Papel")
	fmt.Println("3-Tesoura")
	fmt.Println("4-Lagarto")
	fmt.Println("5-Spock")
	scanner.Scan()
	opcaoPrimeiroJogador, _ := strconv.Atoi(scanner.Text())

	fmt.Println(segundoJogador, ", escolha uma opção: ")
	fmt.Println("1-Pedra")
	fmt.Println("2-Papel")
	fmt.Println("3-Tesoura")
	fmt.Println("4-Lagarto")
	fmt.Println("5-Spock")
	scanner.Scan()
	opcaoSegundoJogador, _ := strconv.Atoi(scanner.Text())

	var resultado = ""
	switch {
	case opcaoPrimeiroJogador == opcaoSegundoJogador || opcaoSegundoJogador == opcaoPrimeiroJogador:
		resultado = "O jogo empatou!"
	case opcaoPrimeiroJogador == 1 && (opcaoSegundoJogador == 3 || opcaoSegundoJogador == 4):
		resultado = fmt.Sprintf("O jogador %s venceu", primeiroJogador)
	case opcaoPrimeiroJogador == 2 && (opcaoSegundoJogador == 1 || opcaoSegundoJogador == 4):
		resultado = fmt.Sprintf("O jogador %s venceu", primeiroJogador)
	case opcaoPrimeiroJogador == 3 && (opcaoSegundoJogador == 2 || opcaoSegundoJogador == 4):
		resultado = fmt.Sprintf("O jogador %s venceu", primeiroJogador)
	case opcaoPrimeiroJogador == 4 && (opcaoSegundoJogador == 2 || opcaoSegundoJogador == 5):
		resultado = fmt.Sprintf("O jogador %s venceu", primeiroJogador)
	case opcaoPrimeiroJogador == 5 && (opcaoSegundoJogador == 1 || opcaoSegundoJogador == 3):
		resultado = fmt.Sprintf("O jogador %s venceu", primeiroJogador)
	default:
		resultado = fmt.Sprintf("O jogador %s venceu", segundoJogador)
	}

	fmt.Println("Resultado:", resultado)
}
