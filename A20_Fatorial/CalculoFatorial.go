package main

import "fmt"

func CalculaFatorial(numero float64) {

	if numero < 0 {
		fmt.Println("ERRO!!! Não é possível calcular fatorial de numero Negativo!!!")
	} else if numero == 0 {
		fmt.Print("FATORIAL CALCULADO: 1")
	}

	Fatorial := 1.00

	for i := 1.00; i <= numero; i++ {
		Fatorial *= i
	}

	if numero == 0 {
		return
	} else if numero < 0 {
		return
	}
	fmt.Print("FATORIAL CALCULADO: ")
	fmt.Println(Fatorial)

}
