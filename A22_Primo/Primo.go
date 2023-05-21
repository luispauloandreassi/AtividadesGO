package main

import (
	"math"
)

func VerificaPrimo(numero int) bool {
	if numero <= 1 {
		return false
	}

	limiteVerificacao := int(math.Sqrt(float64(numero)))

	for i := 2; i <= limiteVerificacao; i++ {
		if numero%i == 0 {
			return false
		}
	}

	return true
}
