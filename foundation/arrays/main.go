package main

import "fmt"

func main() {
	// arrays em go possuem tamanho fixo
	var myArray [3]int

	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3

	// o primeiro valor é o index e o segundo é o valor em sí
	for i, v := range myArray {
		fmt.Printf("O valor do índice é %d e o valor é %d \n", i, v)
	}
}
