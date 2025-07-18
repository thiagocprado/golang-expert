package main

import "github.com/thiagocprado/golang-expert/packaging/scope/math"

func main() {
	newMath := math.NewMath(1, 2)
	println(newMath.Add())

	// Exemplo usando struct e funções exportadas (letra maiúscula)
	exported := math.NewMathExported(10, 20)
	println(exported.AddExported()) // Saída: 30
}
