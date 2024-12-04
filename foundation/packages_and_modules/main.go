package main

import (
	"fmt"
	"foundation/mathematics"

	"github.com/google/uuid"
)

func main() {
	s := mathematics.Sum(10, 20)
	fmt.Printf("Resultado: %v", s)

	// variavel visivel
	println(mathematics.GlobalVar)

	car := mathematics.Car{
		Name: "Chevrolet Onix",
		// a propriedade year não está disponivel
	}
	fmt.Println(car)

	fmt.Println(uuid.New())
}
