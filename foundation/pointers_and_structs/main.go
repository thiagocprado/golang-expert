package main

import "fmt"

type Client struct {
	name        string
	namePointer *string
}

func (c Client) move() {
	// apenas locamente o nome foi alterado, pois estamos trabalhando com uma cópia
	c.name = "Thiago Prado"

	*c.namePointer = "Thiago Prado"
	fmt.Printf("O cliente %s se moveu\n", c.name)
	fmt.Printf("O cliente %s se moveu\n", *c.namePointer)
}

type Account struct {
	balance int
}

// com o ponteiro na struct, estamos trabalhando com os valores reais
func (a *Account) simulate(value int) int {
	a.balance += value

	return a.balance
}

// retorna o endereçamento na memória da conta que está sendo criada
// com isso, em qualquer lugar que passamos a conta e houver alterações, em todos os lugares será alterado
// como se fosse um construtor
func NewAccount() *Account {
	return &Account{}
}

func main() {
	namePointer := "thiago"

	thiago := Client{
		name:        "Thiago",
		namePointer: &namePointer,
	}

	thiago.move()

	fmt.Printf("O valor de name da struct está %s\n", thiago.name)
	fmt.Printf("O valor de namePointer da struct está %s\n", *thiago.namePointer)

	account := Account{
		balance: 100,
	}

	account.simulate(200)
	// o valor muda, porque estamos trabalhando com os valores reais da struct
	println(account.balance)
}
