package main

import "fmt"

type Client struct {
	Name   string
	Age    int
	Active bool
}

// o attachment entre interface e struct é feito automaticamente através dos métodos
// não temos implements em go
// não podemos passar propriedades em interfaces, apenas métodos
type Person interface {
	// também temos assinaturas
	// Activate(id int) bool
	Deactivate()
}

type Company struct {
	Name string
}

func (c Company) Deactivate() {
}

func (c Client) Deactivate() {
	c.Active = false
	fmt.Printf("O cliente %s foi desativado", c.Name)
}

func Deactivation(person Person) {
	person.Deactivate()
}

func main() {
	thiago := Client{
		Name:   "thiago",
		Age:    23,
		Active: true,
	}

	myCompany := Company{
		Name: "myCompany",
	}

	// por debaixo dos panos ambos implementam automaticamente a interface pessoa, através do método deactivate
	Deactivation(thiago)
	Deactivation(myCompany)
}
