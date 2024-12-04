package main

import (
	"fmt"
)

func main() {
	// hash table em go chamamos de map

	salaries := map[string]int{"thiago": 20000, "marcelo": 25000, "murilo": 25000}

	fmt.Println(salaries["thiago"])

	delete(salaries, "murilo")

	salaries["wesley"] = 22000

	salariesTwo := make(map[string]int)
	salariesThree := map[string]int{}

	fmt.Println(salariesTwo)
	fmt.Println(salariesThree)

	for name, salary := range salaries {
		fmt.Printf("O salario de %s é %d\n", name, salary)
	}

	// blank identifier
	for _, salary := range salaries {
		fmt.Printf("O salario é de %d\n", salary)
	}
}
