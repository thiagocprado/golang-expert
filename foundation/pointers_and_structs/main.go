package main

import "fmt"

type Client struct {
	name        string  // Campo do tipo string (valor); alterações locais não afetam o original.
	namePointer *string // Campo do tipo ponteiro para string; alterações afetam diretamente o valor apontado.
}

func (c Client) move() {
	// Aqui, 'c' é uma cópia da struct original.
	// Alterações feitas em 'c.name' afetam apenas essa cópia, e não o objeto original.
	c.name = "Thiago Prado"

	// Já 'c.namePointer' é um ponteiro; ao desreferenciar e alterar,
	// modificamos o valor diretamente na memória, impactando o valor original.
	*c.namePointer = "Thiago Prado"

	// Exibe os valores alterados localmente.
	fmt.Printf("O cliente %s se moveu\n", c.name)
	fmt.Printf("O cliente %s se moveu\n", *c.namePointer)
}

type Account struct {
	balance int // Saldo da conta.
}

// Função "construtor" para criar uma nova conta.
// Retorna o endereço de memória da struct criada.
// Assim, qualquer alteração feita na conta impactará a mesma instância em qualquer lugar do programa.
func NewAccount() *Account {
	return &Account{}
}

// Método com receptor por ponteiro (*Account):
// permite modificar diretamente o valor da struct original.
func (a *Account) simulate(value int) int {
	// Modificamos o saldo real, pois estamos acessando a struct via ponteiro.
	a.balance += value

	return a.balance
}

func main() {
	// Inicializa uma string e cria uma variável ponteiro para ela.
	namePointer := "thiago"

	// Inicializa a struct Client.
	// 'name' recebe uma cópia da string.
	// 'namePointer' recebe o endereço da variável 'namePointer', ou seja, a referência à string.
	thiago := Client{
		name:        "Thiago",
		namePointer: &namePointer,
	}

	// Chama o método 'move'.
	// 'name' será alterado apenas dentro da cópia da struct.
	// '*namePointer' será alterado globalmente, afetando 'namePointer' no escopo main.
	thiago.move()

	// Exibe o valor atual de 'name' na struct (não alterado externamente).
	fmt.Printf("O valor de name da struct está: %s\n", thiago.name)

	// Exibe o valor apontado por 'namePointer', que foi alterado dentro do método.
	fmt.Printf("O valor de namePointer da struct está: %s\n", *thiago.namePointer)

	// Inicializa uma conta com saldo de 100.
	account := Account{
		balance: 100,
	}

	// Chama o método 'simulate', passando um valor de 200.
	// O saldo real será alterado porque o método tem receptor por ponteiro.
	account.simulate(200)

	// Exibe o novo saldo da conta.
	fmt.Printf("O saldo da conta é: %d\n", account.balance)

	var varInt int
	var varIntPointer *int

	fmt.Println(varInt)        // Exibe 0, pois é o valor padrão de um int.
	fmt.Println(varIntPointer) // Exibe nil, pois o ponteiro não aponta para nenhum endereço de memória.
}
