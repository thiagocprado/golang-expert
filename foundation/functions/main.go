package main

import (
	"errors"
	"fmt"
)

func main() {
	val, err := sumErr(50, 10)
	if err != nil {
		fmt.Println(val)
	}

	fmt.Println(sum(10, 23))
}

// podemos retornar mais de um valor
func sum(a, b int) (int, bool) {
	return a + b, a+b >= 50
}

// utilizamos o error como ultimo parametro de resposta
func sumErr(a, b int) (int, error) {
	if a+b >= 50 {
		return a + b, errors.New("A soma é maior que 50")
	}

	return a + b, nil
}
