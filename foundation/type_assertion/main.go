package main

import (
	"fmt"
)

func main() {
	var myVar interface{} = "thiago prado"
	var myVar2 interface{} = 10
	var myVar3 interface{} = []string{"thiago prado"}

	// não mostra o valor pois não sabe o tipo
	println(myVar)

	// mostra o valor pois "definimos" o tipo, chamamos de type assertion
	println(myVar.(string))
	println(myVar2.(int))
	println(myVar3.([]string))

	res, ok := myVar.(int)
	fmt.Printf("O valor de res é %v e o resultado de ok é %v\n", res, ok)

	// se não tivermos o 'ok' o programa irá quebrar para garantir a tipagem forte
	res, ok = myVar2.(int)
	fmt.Printf("O valor de res é %v e o resultado de ok é %v", res, ok)
}
