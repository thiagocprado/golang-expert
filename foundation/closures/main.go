package main

import (
	"fmt"
)

// Closure é uma função que referencia variáveis de seu escopo
//  externo, mesmo após esse escopo ter sido finalizado.

// Em Go, closures são frequentemente funções anônimas, mas
// não precisam necessariamente ser anônimas.

// A característica principal de uma closure não é conter outra
// função dentro dela, mas sim capturar e manter acesso a variáveis do escopo onde foi criada.

func main() {
	total := func() int {
		return sum(10, 23, 21, 40, 64, 53, 78) * 2
	}()

	fmt.Println(total)

	add := adder() // add é agora uma função que "lembra" da variável sum

	fmt.Println(add(1)) // sum = 0 + 1 => 1
	fmt.Println(add(2)) // sum = 1 + 2 => 3
	fmt.Println(add(3)) // sum = 3 + 3 => 6
}

// A função adder retorna uma função (closure) que incrementa e retorna um valor.
func adder() func(int) int {
	sum := 0 // Variável do escopo externo, que será "capturada" pela closure

	// Esta é a closure: ela usa e modifica a variável sum
	return func(x int) int {
		sum += x // A função "lembra" do sum e pode modificá-lo
		return sum
	}
}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}
