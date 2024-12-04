package main

func main() {
	// * declara um pointeiro
	// & obtém o endereço de memória de uma variável

	// os ponteiros permitem trabalhar em um escopo mais amplo, pois pegamos o valor no endereço na memória

	// memória -> endereço -> valor
	// abre uma "caixa" na memória, essa "caixa" tem um endereço e o valor está dentro dela
	// se quiser saber o valor, vai na caixa e pega o valor
	a := 10
	// valor dentro da "caixa"
	println(a)

	// nesse caso estamos trabalhando com uma cópia de 'a' que não referencia o valor real dela, por isso o mesmo não é alterado
	c := a
	c = 40
	println(c)
	println(a)

	// endereço da "caixa"
	println(&a)

	// variavel aponta para um ponteiro que tem um endereço na memória que tem um valor
	// o * se refere ao endereçamento da memória
	var pointer *int = &a
	println(pointer)

	// Desreferenciação - acessamos ou alteramos o valor de uma variável
	// aqui estamos realmente apontando pro valor real de 'a', pois vamos no endereço de memória e fazemos a alteração
	*pointer = 99
	println(a)

	b := &a
	*b = 30
	// o * pega o valor que está na memória, sem ele, nós pegamos o endereço da memória
	println(a)

}
