package main

import (
	"fmt"
	"strconv"
)

func FizzBuzz(valorInicial, numeroFinal int) {
	fmt.Println("Apresentacao:")
	for i := valorInicial; i <= numeroFinal; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			if i == numeroFinal {
				fmt.Print("FizzBuzz")
			} else {
				fmt.Print("FizzBuzz-")
			}

		case i%3 == 0:
			if i == numeroFinal {
				fmt.Print("Fizz")
			} else {
				fmt.Print("Fizz-")
			}
		case i%5 == 0:
			if i == numeroFinal {
				fmt.Print("Buzz")
			} else {
				fmt.Print("Buzz-")
			}
		default:

			if i == numeroFinal {
				concatenated1 := strconv.Itoa(i)
				fmt.Print(concatenated1)
			} else {
				concatenated2 := strconv.Itoa(i) + "-"
				fmt.Print(concatenated2)
			}

		}
	}
}
