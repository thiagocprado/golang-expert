package main

import "fmt"

func main() {
	// Imagine que uma variável é uma caixa com um valor dentro.

	// Essa caixa fica guardada em um endereço na memória do computador.

	// Quando usamos um ponteiro, estamos guardando o endereço dessa caixa, não o valor.

	// Assim, podemos ir até lá e ver ou mudar o valor direto na caixa.

	// * declara um ponteiro
	// & obtém o endereço de memória de uma variável

	// Conceito: a memória pode ser imaginada como uma "caixa".
	// Cada caixa tem um endereço e dentro dela está armazenado um valor.

	// Declarando uma variável comum:
	a := 10
	fmt.Println("1) Valor inicial de a:", a)

	// Criando uma cópia de 'a':
	c := a
	c = 40 // Alteramos a cópia
	fmt.Println("2) Valor de c (cópia de a):", c)
	fmt.Println("3) Valor de a continua o mesmo:", a)

	// Pegando o endereço de memória onde 'a' está guardado:
	fmt.Println("4) Endereço de memória de a:", &a)

	// Criando um ponteiro que aponta para 'a':
	var pointer *int = &a
	fmt.Println("5) Ponteiro que aponta para a:", pointer)

	// Usando o ponteiro para acessar o valor dentro de 'a':
	fmt.Println("6) Valor de a acessado pelo ponteiro:", *pointer)

	// Agora, vamos mudar o valor de 'a' usando o ponteiro:
	*pointer = 99
	fmt.Println("7) Novo valor de a após mudança via ponteiro:", a)

	// Criando outro ponteiro para 'a':
	b := &a
	// Alterando o valor de 'a' usando o segundo ponteiro:
	*b = 30
	fmt.Println("8) Novo valor de a após mudança via ponteiro b:", a)

	// Resumo:
	// - Sem ponteiro: mexemos só na cópia.
	// - Com ponteiro: mexemos direto na variável original.
}
