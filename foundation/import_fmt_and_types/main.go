package main

import "fmt"

// podemos criar tipos personalizados
type ID int

var (
	b bool = true
	i int  = 1
	s string
	f ID = 10
)

func main() {
	fmt.Printf("o tipo de B é %T \n", b)
	fmt.Printf("o tipo de I é %T", i)
	fmt.Printf("o tipo de S é %T", s)
	fmt.Printf("o tipo de F é %T", f)
}
