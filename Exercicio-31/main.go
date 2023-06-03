package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Digite a frase: ")
	frase, _ := reader.ReadString('\n')
	frase = strings.TrimSpace(frase)

	fmt.Print("Digite a quantidade de colunas: ")
	var colunas int
	fmt.Scanln(&colunas)

	linhas := quebrarLinhas(frase, colunas)

	fmt.Println("Frase quebrada em linhas:")
	for _, linha := range linhas {
		fmt.Println(linha)
	}
}

func quebrarLinhas(frase string, colunas int) []string {
	linhas := make([]string, 0)
	palavras := strings.Fields(frase)

	linhaAtual := ""
	espacosDisponiveis := colunas

	for _, palavra := range palavras {
		if len(palavra) > espacosDisponiveis {
			if linhaAtual != "" {
				linhas = append(linhas, linhaAtual)
				linhaAtual = ""
				espacosDisponiveis = colunas
			}

			for len(palavra) > colunas {
				linhas = append(linhas, palavra[:colunas])
				palavra = palavra[colunas:]
			}

			linhaAtual += palavra
			espacosDisponiveis = colunas - len(palavra)
		} else {
			if linhaAtual != "" {
				linhaAtual += " "
				espacosDisponiveis--
			}

			linhaAtual += palavra
			espacosDisponiveis -= len(palavra)
		}

		if espacosDisponiveis == 0 {
			linhas = append(linhas, linhaAtual)
			linhaAtual = ""
			espacosDisponiveis = colunas
		}
	}

	if linhaAtual != "" {
		linhas = append(linhas, linhaAtual)
	}

	return linhas
}
