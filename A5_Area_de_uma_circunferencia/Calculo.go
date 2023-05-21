package main

import "math"

func CalcularCircunferencia(pi, raio float64) float64 {
	return math.Pow(raio, 2) * pi
}
