// todos os arquivos que estão na mesma pasta, tem que possuir o mesmo nome de pacote
// tudo que está no mesmo pacote fica compartilhado entre sí
package main // sempre teremos um package com o nome do diretório, com exceção do main

// constantes não alteram o valor
const a = "Hello World"

// declarada uma variável do tipo booleano
var b bool

// o package main sempre espera a função main
func main() {
	println(a)
	println(b)
}
