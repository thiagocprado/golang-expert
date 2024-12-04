package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(10, 23, 21, 40, 64, 53, 78))
}

// funções variádicas, ou seja, não sabemos quantos argumentos terá, mas sabemos que são todos int
func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}
