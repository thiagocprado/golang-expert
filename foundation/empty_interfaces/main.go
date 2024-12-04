package main

import "fmt"

func main() {
	// interfaces assumem o tipo em cima do valor que lhes é passada
	var x interface{} = 10
	var y interface{} = "hello world!"

	showType(x)
	showType(y)
}

func showType(t interface{}) {
	fmt.Printf("O tipo da variável é %T e o valor é %v\n", t, t)
}
