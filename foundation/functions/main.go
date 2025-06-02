package main

import (
	"errors"
	"fmt"
)

func main() {
	val, err := sumErr(50, 10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)

	val, err = sumErr(20, 10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)

	fmt.Println(sum(10, 23))

	sumNoReturn(10, 20)
}

// podemos retornar mais de um valor
func sum(a, b int) (int, bool) {
	return a + b, a+b >= 50
}

// podemos não retornar valor
func sumNoReturn(a, b int) {
	sum := a + b
	fmt.Println(sum)
}

// utilizamos o error como ultimo parametro de resposta
func sumErr(a, b int) (int, error) {
	if a+b >= 50 {
		return a + b, errors.New("A soma é maior que 50")
	}

	return a + b, nil
}
