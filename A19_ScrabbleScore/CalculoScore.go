package main

import (
	"fmt"
	"strings"
)

func CalculoScore(palavra string) {

	PalavraMaiuscula := strings.ToUpper(palavra)
	var Score = 0
	//fmt.Println(Score)
	for _, letra := range PalavraMaiuscula {

		switch letra {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			Score++
		case 'D', 'G':
			Score = Score + 2
		case 'B', 'C', 'M', 'P':
			Score = Score + 3

		case 'F', 'H', 'V', 'W', 'Y':
			Score = Score + 4

		case 'K':
			Score = Score + 5

		case 'J', 'X':
			Score = Score + 8

		case 'Q', 'Z':
			Score = Score + 10
		default:
			Score = Score + 0
		}

	}

	fmt.Print("SCORE CALCULADO: ")
	fmt.Print(Score)

}
