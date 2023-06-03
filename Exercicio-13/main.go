package main

import "fmt"

func main() {
	var distancia, tempo float64

	fmt.Print("Digite a distância percorrida (em quilômetros): ")
	fmt.Scanln(&distancia)

	fmt.Print("Digite o tempo decorrido (em minutos): ")
	fmt.Scanln(&tempo)

	velocidade := (distancia * 1000) / (tempo * 60)

	fmt.Printf("A velocidade do projétil é %.2f metros por segundo\n", velocidade)
}
