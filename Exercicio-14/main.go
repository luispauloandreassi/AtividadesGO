package main

import (
	"fmt"
)

func main() {
	var votosA, votosB, votosC, votosNulos, votosBranco int

	fmt.Print("Digite a quantidade de votos válidos para o candidato A: ")
	fmt.Scanln(&votosA)

	fmt.Print("Digite a quantidade de votos válidos para o candidato B: ")
	fmt.Scanln(&votosB)

	fmt.Print("Digite a quantidade de votos válidos para o candidato C: ")
	fmt.Scanln(&votosC)

	fmt.Print("Digite a quantidade de votos nulos: ")
	fmt.Scanln(&votosNulos)

	fmt.Print("Digite a quantidade de votos em branco: ")
	fmt.Scanln(&votosBranco)

	totalEleitores := votosA + votosB + votosC + votosNulos + votosBranco

	percentVotosValidos := float64(votosA+votosB+votosC) / float64(totalEleitores) * 100
	percentVotosA := float64(votosA) / float64(totalEleitores) * 100
	percentVotosB := float64(votosB) / float64(totalEleitores) * 100
	percentVotosC := float64(votosC) / float64(totalEleitores) * 100
	percentVotosNulos := float64(votosNulos) / float64(totalEleitores) * 100
	percentVotosBranco := float64(votosBranco) / float64(totalEleitores) * 100

	fmt.Println("Resultado da eleição sindical:")
	fmt.Printf("Total de eleitores: %d\n", totalEleitores)
	fmt.Printf("Percentual de votos válidos em relação aos eleitores: %.2f%%\n", percentVotosValidos)
	fmt.Printf("Percentual de votos do candidato A em relação aos eleitores: %.2f%%\n", percentVotosA)
	fmt.Printf("Percentual de votos do candidato B em relação aos eleitores: %.2f%%\n", percentVotosB)
	fmt.Printf("Percentual de votos do candidato C em relação aos eleitores: %.2f%%\n", percentVotosC)
	fmt.Printf("Percentual de votos nulos em relação aos eleitores: %.2f%%\n", percentVotosNulos)
	fmt.Printf("Percentual de votos em branco em relação aos eleitores: %.2f%%\n", percentVotosBranco)
}
