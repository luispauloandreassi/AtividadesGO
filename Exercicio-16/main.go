package main

import "fmt"

func main() {
	inicio := 1
	fim := 100

	soma := 0

	for i := inicio; i <= fim; i++ {
		soma += i
	}

	fmt.Printf("Intervalo: %d a %d\n", inicio, fim)
	fmt.Printf("Soma dos valores: %d\n", soma)
}
