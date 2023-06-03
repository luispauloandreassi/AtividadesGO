package main

import (
	"fmt"
	"sort"
)

func main() {
	numeros := make([]int, 12)

	for i := 0; i < 12; i++ {
		fmt.Printf("Digite o elemento %d: ", i+1)
		fmt.Scanln(&numeros[i])
	}

	sort.Sort(sort.Reverse(sort.IntSlice(numeros)))

	fmt.Println("Elementos ordenados em ordem decrescente:")
	for _, num := range numeros {
		fmt.Println(num)
	}
}
