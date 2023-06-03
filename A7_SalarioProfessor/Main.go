package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	fmt.Println("Informe as horas trabalhadas no mês: ")
	scanner.Scan()
	horasTrabalhadas, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Println("Informe o valor da hora-aula: ")
	scanner.Scan()
	horaAula, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Println("Informe o percentual de desconto: ")
	scanner.Scan()
	percentualDesconto, _ := strconv.ParseFloat(scanner.Text(), 64)

	salarioBruto := horasTrabalhadas * horaAula
	totalDesconto := salarioBruto * (percentualDesconto / 100)
	salarioLiquido := salarioBruto - totalDesconto

	fmt.Println("Salário bruto: ", salarioBruto, " | Salário líquido: ", salarioLiquido)
}
