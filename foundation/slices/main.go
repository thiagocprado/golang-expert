package main

import "fmt"

func main() {
	// slices em go são arrays por debaixo dos panos, porém com tamanho variado
	// slice possuí um ponteiro que aponta pro array
	// um tamanho para saber até onde ir
	// e a capacidade pra saber quanta informação ele pode armazenar
	mySlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Printf("len=%d cap=%d %v\n", len(mySlice), cap(mySlice), mySlice)

	// os : é basicamente criar um slice de um slice, é um corte
	fmt.Printf("len=%d cap=%d %v\n", len(mySlice[:0]), cap(mySlice[:0]), mySlice[:0])

	fmt.Printf("len=%d cap=%d %v\n", len(mySlice[:4]), cap(mySlice[:4]), mySlice[:4])

	fmt.Printf("len=%d cap=%d %v\n", len(mySlice[2:]), cap(mySlice[2:]), mySlice[2:])

	// por debaixo dos panos o go dobra o tamanho do valor do array
	mySlice = append(mySlice, 10)
	fmt.Printf("len=%d cap=%d %v\n", len(mySlice[2:]), cap(mySlice[2:]), mySlice[2:])
}
