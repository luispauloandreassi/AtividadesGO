package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Bem vindo ao programa JokenPo: ")

	fmt.Println("Informe o Numero das Opcoes:")
	fmt.Println("1 - Pedra")
	fmt.Println("2 - Papel")
	fmt.Println("3 - Tesoura")

	fmt.Print("JOGADOR 1: ")
	scanner.Scan()
	jogador1 := scanner.Text()
	opcaoJogador1 := ConverteOpcoes(jogador1)
	if opcaoJogador1 == "Invalido" {
		fmt.Println("Opcao Invalida!!!")
		fmt.Println("O programa esta sendo finalizado...")
		os.Exit(1)
	}

	fmt.Print("JOGADOR 2: ")
	scanner.Scan()
	jogador2 := scanner.Text()
	opcaoJogador2 := ConverteOpcoes(jogador2)
	if opcaoJogador2 == "Invalido" {
		fmt.Println("Opcao Invalida!!!")
		fmt.Println("O programa esta sendo finalizado...")
		os.Exit(1)
	}

	fmt.Println("==================== ")
	fmt.Print("Jogador 1: ")
	fmt.Println(ConverteOpcoes(jogador1))
	fmt.Print("Jogador 2: ")
	fmt.Println(ConverteOpcoes(jogador2))

	JuizJoKenPon(ConverteOpcoes(jogador1), ConverteOpcoes(jogador2))

}
