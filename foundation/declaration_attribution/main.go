package main

// constantes não alteram o valor
const a = "Hello World"

// declarada uma variável do tipo booleano
var b bool

// sempre terá um valor default, independente do tipo
// variáveis possuem escopo, nesse caso, estamos em um escopo global
// variáveis possuem tipos, e os mesmos não podem ser alterados após declarados
var (
	// podemos declarar e já atribuir um valor default
	i int = 1
	s string
	f float64
)

func main() {
	// temos também a inferência de tipos
	// essa declaração é feita apenas a primeira vez
	// short-hand (atalho)
	str := "string"
	integer := 10

	// após declarada, não precisamos mais dos :
	str = "string atribuida"

	// aqui estamos em um escopo local
	var a string

	println(a)
	println(b)
	println(i)
	println(s)
	println(str)
	println(integer)
}
