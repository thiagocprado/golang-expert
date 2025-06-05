package mathematics

import "fmt"

func Sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	println(localVar)

	car := Car{
		year: 2019,
		Name: "Chevrolet Onix",
	}
	fmt.Println(car)

	newBike := bike{
		name: "Yamaha",
		year: 2020,
	}
	fmt.Println(newBike)

	return total
}

// o golang é case sensitive quando estamos trabalhando com escopo local e global
// ou seja, se temos variaveis com letra minuscula, ela é considerada de escopo local
// então não vai ser visivel fora do pacote onde foi decladara
var localVar string = "apenas no escopo local"

// agora quando temos uma letra maiuscula, nós temos a visão fora do escopo local
var GlobalVar string

// o mesmo vale para nomes de funções e structs
type Car struct {
	Name string
	year int
}

type bike struct {
	name string
	year int
}
