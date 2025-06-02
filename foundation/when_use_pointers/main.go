package main

// A função sum recebe:
// - 'a' como um valor, ou seja, uma cópia da variável original.
// - 'b' como um ponteiro para um int, ou seja, a referência para a variável original.
// Modificações em 'a' não afetam a variável original, mas modificações via 'b' afetam diretamente.
func sum(a int, b *int) int {
	a = 50  // Modifica apenas a cópia local de 'a'; não afeta a variável original.
	*b = 30 // Modifica o valor na memória apontada por 'b'; afeta a variável original.

	return a + *b // Retorna a soma da cópia de 'a' com o valor apontado por 'b'.
}

func main() {
	// Declaramos duas variáveis inteiras
	myVar := 10
	myVar2 := 20

	// Chamamos a função sum passando:
	// - 'myVar' como valor (cópia)
	// - '&myVar2' como ponteiro (endereço de memória)
	sum(myVar, &myVar2)

	// Como 'myVar' foi passado por valor, a função sum não alterou seu conteúdo.
	println(myVar) // Saída: 10

	// Como 'myVar2' foi passado por referência (ponteiro),
	// a função sum alterou diretamente o seu conteúdo.
	println(myVar2) // Saída: 30

	// Resumo:
	// - Use passagem por valor quando não quiser modificar a variável original.
	// - Use ponteiros quando quiser que a função possa modificar o valor da variável.
}
