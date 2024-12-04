package main

import (
	"fmt"
)

func main() {
	// closures funcionam basicamente como funções anônimas
	func() {
	}()

	total := func() int {
		return sum(10, 23, 21, 40, 64, 53, 78) * 2
	}()

	fmt.Println(total)
}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}
