# Go Expert - Full Cycle

> 📚 **Projeto do curso Full Cycle - Go Expert**
> ✅ **Propósito**: consolidar os principais conceitos da linguagem **Go** de forma clara e objetiva.

---

## 📌 Sobre a Linguagem Go

### ✅ O que é Go?

* **Go** é uma linguagem de programação **open-source**.

* Criada para aumentar a **produtividade** dos desenvolvedores.

* Focada em ser:

  * **Expressiva**
  * **Concisa**
  * **Limpa**
  * **Eficiente**

* Projetada para aproveitar ao máximo:

  * Arquiteturas **multicore** (múltiplos núcleos de processamento).
  * Sistemas de **rede**.

---

### ✅ Principais Características

* ✅ **Compilação rápida** + **garbage collector** automático.
* ✅ **Estática** e **compilada**, com sintaxe que lembra linguagens **dinâmicas**.
* ✅ Gera um **único arquivo binário** → facilitando o deploy.
* ✅ **Retrocompatibilidade**: desde a versão **1.0** (2012).
  Não há necessidade de se preocupar com quebras em versões futuras.

---

### ✅ Origem da Linguagem

* Criada no **Google** por:

  * **Rob Pike** (Unix, UTF-8)
  * **Robert Griesemer** (V8)
  * **Ken Thompson** (Unix, UTF-8)

* Desenvolvimento iniciado em **2007**.

* Primeira versão oficial: **2012**.

* A partir da versão **1.5**, o compilador Go passou a ser escrito em **Go**.

---

### ✅ O que Go **não é**?

* ❌ Não é uma linguagem que resolverá **todos os problemas**.
* ❌ Não é **dinâmica**.
* ❌ Não é **interpretada**.
* ❌ Não possui uma infinidade de recursos → foca na **simplicidade**.

---

### ✅ Motivação para a criação

Superar limitações de linguagens populares da época:

| Linguagem  | Limitação                                           |
| ---------- | --------------------------------------------------- |
| **Python** | Problemas de **desempenho**                         |
| **C/C++**  | **Complexidade** e tempo elevado de desenvolvimento |
| **Java**   | Crescente **complexidade** e **verbosidade**        |

* Linguagens tradicionais não nasceram focadas em **concorrência** e **multithreading** → Go foi **projetada para isso** desde o início.

---

## ⚡ Multithreading e Concorrência

* Destaque: **modelo de concorrência** baseado em **goroutines** e **channels**.

### ✅ Exemplo simples de goroutine:

```go
package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello, Go!")
}

func main() {
	go sayHello()
	time.Sleep(1 * time.Second)
}
```

> 📝 A palavra-chave `go` inicia uma nova **goroutine** (thread leve gerenciada pela runtime do Go).

---

## ✅ Pontos Fortes da Linguagem Go

* ✅ **Simplicidade** → código fácil de entender e manter.
* ✅ **Ferramentas nativas**:

  * **Testes** (`testing`)
  * **Profiling** (`pprof`)
  * **Detecção de race conditions** (`-race`)
* ✅ **Deploy** simplificado → apenas **um arquivo binário**.
* ✅ **Baixa curva de aprendizado**.

---

## ⚠️ Pontos de Atenção

* Go possui um **jeito próprio** de programar.
* Cuidado para não replicar padrões de outras linguagens que **não fazem sentido** no contexto de Go.

Boas práticas:

* ✅ Prefira **composição** ao invés de **herança**.
* ✅ Use **interfaces pequenas** e coesas.
* ✅ Mantenha o código **simples** e **explícito**.

---

## ✅ Como funciona o runtime do Go?

```
Código GO + Nosso Código = Arquivo binário
 RUNTIME
   src/
```

* O **arquivo binário** pode ser compilado para diferentes **sistemas operacionais** e **arquiteturas**.

### ✅ Variáveis importantes:

* **GOOS** → define o sistema operacional (Ex.: `linux`, `windows`).
* **GOARCH** → define a arquitetura do processador (Ex.: `amd64`, `arm64`).
* **GOPATH** → local onde o Go armazena os pacotes (desde Go 1.11, com modules, é menos usado).

---

## ✅ Programação Orientada a Objetos em Go

### ⚡ Go não possui **classes**!

Mas:

* ✅ Possui **tipos** (`structs`).
* ✅ Podemos associar **métodos** aos tipos.
* ✅ Go possui **interfaces** para garantir a **assinatura de métodos**.

---

### ✅ Modificadores de acesso

* Não há `public`, `private`, `protected`.
* ✅ **Convenção**:

  * Letra **maiúscula** → **exportado** (público).
  * Letra **minúscula** → **não exportado** (privado).

```go
type Pessoa struct {
	Nome string // Exportado
	idade int   // Não exportado
}
```

---

### ✅ Herança e Polimorfismo

* Go **não** possui **herança** tradicional.
* ✅ Go usa **composição** para alcançar comportamento **similar**.

```go
type Animal struct {
	Nome string
}

func (a Animal) Falar() string {
	return "O animal faz um som"
}

type Cachorro struct {
	Animal // Composição
}

func (c Cachorro) Falar() string {
	return "Au au!"
}
```

---

### ✅ Polimorfismo

* Alcançado através de **interfaces** → **polimorfismo implícito**.

```go
type Falante interface {
	Falar() string
}

func FazerFalar(f Falante) {
	fmt.Println(f.Falar())
}

func main() {
	c := Cachorro{Animal{"Rex"}}
	FazerFalar(c)
}
```

---

## ✅ Paradigmas predominantes

* ✅ **Imperativo** → descrevemos "como" executar.
* ✅ **Concorrente** → suporte nativo via goroutines e channels.

---

## ✅ Tipagem em Go

* ✅ **Tipagem estrutural** → se um tipo **possui os métodos** requeridos por uma interface, ele **implementa** a interface, **sem necessidade de declaração explícita**.
* ✅ Suporta **duck typing** → "se anda como um pato...".

---

## ✅ Compilação cruzada

* Go permite compilar para **outras plataformas** de forma simples:

```bash
GOOS=linux GOARCH=amd64 go build -o app-linux
```

---

## ✅ Complementos importantes

### 📌 Go Modules

* Desde a versão **1.11**, o gerenciamento de dependências com **Go Modules** (`go mod init`, `go mod tidy`) substitui a necessidade do `GOPATH`.

---

### 📌 Formatação e padrões

* Go possui **formatação automática** com `gofmt`.
* **Código idiomático** é altamente valorizado → leia e siga o **Effective Go** e os **Go Proverbs**.

---

### 📌 Comunidade e uso

* Muito usada em:

  * Desenvolvimento de **microserviços**.
  * **Ferramentas de rede** e **infraestrutura**.
  * **Cloud computing** → Docker, Kubernetes, Terraform são exemplos de ferramentas escritas em Go.

---

## ✅ Referências oficiais

* [Go.dev](https://go.dev)
* [Effective Go](https://go.dev/doc/effective_go)
* [Go Proverbs](https://go-proverbs.github.io/)

---

## ✅ Resumo final

✅ Go é simples, poderosa e eficiente.
✅ Focada em produtividade, desempenho e facilidade de manutenção.
✅ Ideal para aplicações modernas, especialmente na **nuvem**.


// TO-DO
- pesquisar sobre as ferramentas do go para vscode e adicionar um resumo delas aqui 

- pesquisar o porque generecis podem ser um problema

- colocar também sobre go.mod e go.sum