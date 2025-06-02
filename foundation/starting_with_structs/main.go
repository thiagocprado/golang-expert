package main

import "fmt"

// golang não é e não é uma linguagem orientada a objetos
// ela possuí conceitos semelhantes a POO
// mas ela possuí o jeito go de programar

// a struct é semelhante ao objeto
type Client struct {
	Name   string
	Age    int
	Active bool
}

// assim nós podemos atribuir métodos a uma struct, como se fosse os métodos de uma classe
func (c Client) Deactivate() {
	c.Active = false
	fmt.Printf("O cliente %s foi desativado / status %v", c.Name, c.Active)
}

func main() {
	thiago := Client{
		Name:   "thiago",
		Age:    23,
		Active: true,
	}

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %v", thiago.Name, thiago.Age, thiago.Active)
	thiago.Deactivate()
}
