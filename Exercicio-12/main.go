package main

import "fmt"

func main() {
	var num1, num2, num3 float64

	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&num1)

	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&num2)

	fmt.Print("Digite o terceiro número: ")
	fmt.Scanln(&num3)

	soma := num1 + num2 + num3
	quadradoDaSoma := soma * soma

	fmt.Printf("O quadrado da soma dos três valores é: %.2f\n", quadradoDaSoma)
}
