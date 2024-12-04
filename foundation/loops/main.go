package main

import "fmt"

func main() {
	// o go trabalha apenas com o for
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	numbers := []string{"one", "two", "three"}

	// index, valor
	for i, v := range numbers {
		fmt.Println(i, v)
	}

	// se assemelha com o do while
	j := 0
	for j < 10 {
		fmt.Println(j)
		j++
	}

	// loop infinito
	for {
		fmt.Println("loop")
	}
}
