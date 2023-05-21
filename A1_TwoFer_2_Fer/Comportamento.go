package main

import "fmt"

func twofer(nome string) {
	if nome == "" {
		nome = "você"
	}
	fmt.Printf("Um para mim, um para %s.\n", nome)
}
