package main

// recebom o 'a' que é uma cópia e o 'b' que é o valor da memória em si
func sum(a int, b *int) int {
	a = 50
	*b = 30

	return a + *b
}

func main() {
	// aqui temos apenas uma cópia dos valores
	myVar := 10
	myVar2 := 20

	sum(myVar, &myVar2)

	// como passamos apenas uma cópia, o valor de myVar não será alterado
	println(myVar)

	// quando passamos o endereço da memória, nós alteramos o valor real da variável
	println(myVar2)

	// não usamos ponteiro quando queremos passar apenas uma cópia
	// usamos pointerios quando queremos trabalhar com valores mutáveis
}
