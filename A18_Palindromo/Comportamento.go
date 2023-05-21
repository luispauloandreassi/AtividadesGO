package main

import "fmt"

func Invertedados(dados string) string {
	arrayFrase := []rune(dados)
	inicio := 0
	fim := len(arrayFrase) - 1

	for inicio < fim {
		arrayFrase[inicio], arrayFrase[fim] = arrayFrase[fim], arrayFrase[inicio]
		inicio++
		fim--
	}
	return string(arrayFrase)
}

func VerificaPalindromo(dados string, dadosInvertido string) {

	if dados == dadosInvertido {
		fmt.Println("RESULTADO: É PALINDROMO!!!")
	} else {
		fmt.Println("RESULDADO: NÃO É PALINDROMO!!!")
	}

}
