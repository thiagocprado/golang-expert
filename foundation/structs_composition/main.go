package main

import "fmt"

// go não possuí herança, mas podemos compor structs
// chamamos de composição

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
	// structs podem servir como tipo também
	Conta Account
}

func main() {
	thiago := Client{
		Name:   "thiago",
		Age:    23,
		Active: true,
	}

	// mesmo resultado
	thiago.City = "Franca"
	thiago.Address.State = "SP"

	thiago.Conta.Number = 123456

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %v", thiago.Name, thiago.Age, thiago.Active)
}
