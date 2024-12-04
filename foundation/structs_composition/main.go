package main

import "fmt"

type Address struct {
	Street string
	Number int
	City   string
	State  string
}

type Account struct {
	Number  int
	Balance float64
}

// podemos encadear structs, fazendo uma composição delas, com isso conseguimos separar as informações
type Client struct {
	Name   string
	Age    int
	Active bool
	// Se tiver o mesmo nome da struct, podemos simplificar a escrita
	Address
	Account Account
}

func main() {
	thiago := Client{
		Name:   "thiago",
		Age:    23,
		Active: true,
	}

	thiago.Address.City = "Franca"

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %v", thiago.Name, thiago.Age, thiago.Active)
}
