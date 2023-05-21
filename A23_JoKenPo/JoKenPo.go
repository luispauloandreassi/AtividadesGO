package main

import "fmt"

func JuizJoKenPon(jogador1 string, jogador2 string) {

	switch {
	case jogador1 == jogador2:
		fmt.Println("QUE PENA! OCORREU EMPATE!")
		fmt.Println("O programa esta sendo finalizado...")
	case jogador1 == "Pedra" && jogador2 == "Tesoura":
		fmt.Println("JOGADOR 1 VENCEU!!!")
	case jogador1 == "Papel" && jogador2 == "Pedra":
		fmt.Println("JOGADOR 1 VENCEU!!!")
	case jogador1 == "Tesoura" && jogador2 == "Papel":
		fmt.Println("JOGADOR 1 VENCEU!!!")
	default:
		fmt.Println("JOGADOR 2 VENCEU!!!")
	}

}

func ConverteOpcoes(opcao string) string {

	opcaoFinal := " "

	if opcao == "1" {
		opcaoFinal = "Pedra"
	}
	if opcao == "2" {
		opcaoFinal = "Papel"
	}
	if opcao == "3" {
		opcaoFinal = "Tesoura"
	}
	if opcao != "1" && opcao != "2" && opcao != "3" {
		opcaoFinal = "Invalido"
	}

	return opcaoFinal
}
