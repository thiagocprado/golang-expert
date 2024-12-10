package main

import "net/http"

func main() {
	// Criamos um multiplexer (mux), que é responsável por rotear as solicitações HTTP
	// para diferentes handlers com base nos caminhos (URLs).
	mux := http.NewServeMux()

	// Registramos o handler "HomeHandler" para o caminho raiz ("/").
	mux.HandleFunc("/", HomeHandler) // HandleFunc associa funções ao mux.

	// Registramos o handler "blog", que é uma struct que implementa a interface http.Handler.
	// Aqui, o caminho "/blog" será tratado pela struct "blog".
	mux.Handle("/blog", blog{title: "My Blog!"})

	// Inicia o servidor HTTP na porta 8080 e usa o mux para gerenciar as solicitações.
	// Ao invés de usarmos o mux goblal do go, criamos o nosso próprio para termos mais controle sobre ele
	// O mux padrão pode permitir que outros pacotes injetem coisas que não queremos
	// Também não nos permite criar servidores diferentes
	http.ListenAndServe(":8080", mux)
}

// HomeHandler é uma função que será chamada sempre que uma solicitação for feita ao caminho "/".
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Escreve a resposta "Hello World!" no corpo da resposta HTTP.
	w.Write([]byte("Hello World!"))
}

// blog é uma struct que define um blog com um título.
type blog struct {
	title string
}

// ServeHTTP é um método da struct "blog" que implementa a interface http.Handler.
// Ele será chamado automaticamente quando houver uma solicitação para o caminho "/blog".
func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Escreve o título do blog no corpo da resposta HTTP.
	w.Write([]byte(b.title))
}

// essa abordagem permite uma customização maior, porém teriamos diversas structs
