package main

import "fmt"

// O vínculo (attachment) entre interface e struct é feito automaticamente através dos métodos.
// Não há 'implements' ou 'extends' como em outras linguagens.
// Interfaces não podem ter propriedades, apenas métodos. Elas definem comportamentos.

// Exemplo de interface com um método
type Person interface {
	Deactivate()
}

// Outra interface, com uma assinatura diferente
type Document interface {
	GetCNPJ() string
}

// Struct que terá um método para implementar a interface Document
type Company struct {
	Name string
	CNPJ string
}

// Como Company possui o método GetCNPJ, ela implementa automaticamente a interface Document
func (c Company) GetCNPJ() string {
	return c.CNPJ
}

// Struct Client, que implementará a interface Person
type Client struct {
	Name   string
	Age    int
	Active bool
}

// Como Client possui o método Deactivate, ela implementa automaticamente a interface Person
func (c Client) Deactivate() {
	c.Active = false
	fmt.Printf("O cliente %s foi desativado / status: %v\n", c.Name, c.Active)
}

// Função que aceita qualquer valor que implemente a interface Person
func Deactivation(person Person) {
	person.Deactivate()
}

// Função que aceita qualquer valor que implemente a interface Document
func GetDocument(d Document) {
	fmt.Println(d.GetCNPJ())
}

func main() {
	thiago := Client{
		Name:   "thiago",
		Age:    23,
		Active: true,
	}

	company := Company{
		Name: "myCompany",
		CNPJ: "12345678000100",
	}

	// thiago implementa Person via método Deactivate
	Deactivation(thiago) // Passar como ponteiro, pois Deactivate altera o estado

	// company implementa Document via método GetCNPJ
	GetDocument(company)
}
