package main

func CalcularSalario(salario, percentual float64) float64 {
	return salario + (salario * percentual / 100)
}
