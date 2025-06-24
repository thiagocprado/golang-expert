package main

import "net/http"

func main() {
	// Criamos um novo *ServeMux*, que é um roteador de requisições HTTP.
	// Ele permite associar diferentes caminhos (URLs) a diferentes handlers.
	// Isso oferece mais controle do que usar o mux padrão do pacote http.
	mux := http.NewServeMux()

	// Associa a função HomeHandler ao caminho raiz ("/").
	// Sempre que uma requisição for feita a "/", essa função será executada.
	mux.HandleFunc("/", HomeHandler)

	// Associa o caminho "/blog" a uma instância da struct "blog",
	// que implementa a interface http.Handler.
	// Isso demonstra como structs podem ser usadas como manipuladores personalizados.
	mux.Handle("/blog", blog{title: "My Blog!"})

	// Inicia o servidor HTTP na porta 8080 utilizando o mux criado como roteador.
	// Ao usar um ServeMux próprio (em vez do padrão), ganhamos isolamento,
	// mais controle sobre as rotas e podemos criar múltiplos servidores se necessário.
	http.ListenAndServe(":8080", mux)
}

// HomeHandler trata requisições feitas ao caminho "/".
// Ele envia uma resposta simples com o texto "Hello World!".
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

// blog é uma struct que representa um blog com um título.
type blog struct {
	title string
}

// ServeHTTP implementa a interface http.Handler para a struct blog.
// Esse método será chamado sempre que uma requisição for feita ao caminho "/blog".
func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}

// Essa abordagem, utilizando structs como handlers, permite maior flexibilidade
// e encapsulamento. No entanto, pode resultar em muitas structs se o número de
// rotas e comportamentos personalizados crescer muito.
