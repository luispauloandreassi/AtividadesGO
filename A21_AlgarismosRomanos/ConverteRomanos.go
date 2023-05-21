package main

import "regexp"

func ConverteRomano(algarismos string) int {

	deParaRomanos := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	valoInteiro := 0

	for i := 0; i < len(algarismos); i++ {
		valor := deParaRomanos[algarismos[i]]
		if i+1 < len(algarismos) && valor < deParaRomanos[algarismos[i+1]] {
			valoInteiro -= valor
		} else {
			valoInteiro += valor
		}
	}

	return valoInteiro
}

func validarAlgarismosRomanos(algarismos string) bool {
	correto := regexp.MustCompile(`^(M{0,3})(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$`)
	return correto.MatchString(algarismos)
}
